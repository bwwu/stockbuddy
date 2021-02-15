package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"stockbuddy/analysis/detectors/macdx"
	"stockbuddy/analysis/detectors/smax"
	"stockbuddy/analysis/detectors/swingrejection"
	"stockbuddy/analysis/insight"
	"stockbuddy/fileio"
	pb "stockbuddy/protos/quote_go_proto"
	"stockbuddy/smtp"
	"strings"
	"time"

	"google.golang.org/grpc"
)

var (
	flagNomail    = flag.Bool("nomail", false, "Set to true to disable sending email reports.")
	flagWatchlist = flag.String("use_watchlist", "watchlists/default.txt", "Path to txt file with a list of stocks to track.")
	flagMailList  = flag.String("mail_to", "", "Comma separated list of email addresses to whom results will be mailed.")
	emailRE       = regexp.MustCompile("\\w+@\\w+\\.\\w+")
)

func main() {
	flag.Parse()

	watchlist, err := fileio.ReadLines(*flagWatchlist)
	if err != nil {
		log.Fatal(err.Error())
	}

	mailList, err := parseEmailsFromList(*flagMailList)
	if err != nil {
		log.Fatal(err)
	}

	emailPassword := os.Getenv("STOCKBUDDY_PASSWORD")
	if emailPassword == "" && !*flagNomail {
		log.Fatal("main: password unset. Set email credentials on env $STOCKBUDDY_PASSWORD.")
	}

	t1 := time.Now()
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
	client := pb.NewQuoteServiceClient(conn)

	summaries := process(client, watchlist)
	if len(summaries) > 0 {
		for _, summ := range summaries {
			log.Printf(`main: %d indicator(s) found for "%s".`, len(summ.Indicators), summ.Symbol)
		}
		msgBody := insight.TableFormat(summaries)
		log.Printf("main: analysis took %d ms", time.Now().Sub(t1).Milliseconds())

		if !*flagNomail {
			mail(emailPassword, msgBody, mailList)
		}
	}
}

func mail(password, content string, recipients []string) {
	subject := "Reversal Trends Detected"

	body := "<p>Reversal trends have been detected for the following stocks:</p>\n" + content

	email := smtp.Email{body, subject, recipients}
	email.Send(password)
}

// Given a comma-separated-list of emails given by a flag value, return a list of validated email
// addresses.
func parseEmailsFromList(raw string) ([]string, error) {
	if *flagNomail {
		return []string{}, nil
	}
	errPrefix := "main::parseEmailFromList():"
	result := strings.Split(raw, ",")

	if len(result) == 0 {
		return nil, fmt.Errorf("%s empty email list", errPrefix)
	}

	// Validate email addresses.
	for _, email := range result {
		if !emailRE.MatchString(email) {
			return nil, fmt.Errorf(`%s invalid email "%s"`, errPrefix, email)
		}
	}
	return result, nil
}

func process(client pb.QuoteServiceClient, stocks []string) []*insight.AnalyzerSummary {
	// Instantiate all of the detectors to run.
	detectors := make([]insight.Detector, 0, 3)
	if smaDetec, err := smax.NewSimpleMovingAverageDetector(12, 48); err != nil {
		log.Fatal(err)
	} else {
		detectors = append(detectors, smaDetec)
	}
	if macdDetec, err := macdx.NewMACDDetector(12, 26, 9); err != nil {
		log.Fatal(err)
	} else {
		detectors = append(detectors, macdDetec)
	}
	detectors = append(detectors, swingrejection.NewSwingRejectionDetector(30, 14))

	// Spawn goroutine to run analyzer over all detectors, one per stock.
	summaryc := make(chan *insight.AnalyzerSummary)
	defer close(summaryc)

	for _, symbol := range stocks {
		go func(s string) {
			analyzer := insight.NewAnalyzer(client, detectors...)
			indicators := analyzer.Analyze(context.Background(), s)
			if indicators == nil {
				summaryc <- nil
			} else {
				summaryc <- &insight.AnalyzerSummary{s, indicators}
			}
		}(symbol)
	}

	// Collect result from each analyzer, filter out nil.
	result := make([]*insight.AnalyzerSummary, 0)
	for i := 0; i < len(stocks); i++ {
		summary := <-summaryc
		if summary != nil {
			result = append(result, summary)
		}
	}
	return result
}
