package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"stockbuddy/analysis/detectors/pricedelta"
	"stockbuddy/analysis/insight"
	"stockbuddy/fileio"
	quotepb "stockbuddy/protos/quote_go_proto"
	"stockbuddy/smtp"
	"strings"
	"time"

	"google.golang.org/grpc"
)

var (
	flagNomail    = flag.Bool("nomail", false, "Set to true to disable sending email reports.")
	flagWatchlist = flag.String("use_watchlist", "watchlists/daily_price_watch.txt", "Path to txt file with a list of stocks to track.")
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

	client := quotepb.NewQuoteServiceClient(conn)
	summaries := process(client, watchlist)
	if len(summaries) > 0 {
		for _, summ := range summaries {
			log.Printf(`dailyprices: %d indicator(s) found for "%s".`, len(summ.Indicators), summ.Symbol)
		}
		msgBody := insight.FormatByIndicator(summaries)
		log.Printf("dailyprices: analysis took %d ms", time.Now().Sub(t1).Milliseconds())

		if !*flagNomail {
			mail(emailPassword, msgBody, mailList)
		}
	}
}

func mail(password, content string, recipients []string) {
	subject := "Daily price summary"

	body := "<p>Your watchlist stocks. Note that the days indicates the number of <strong>trading days</strong></p>\n" + content

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

func process(client quotepb.QuoteServiceClient, stocks []string) []*insight.AnalyzerSummary {
	detectors := []insight.Detector{pricedelta.NewDefaultDetector()}
	d, _ := pricedelta.NewDetector(5)
	detectors = append(detectors, d)
	d, _ = pricedelta.NewDetector(30)
	detectors = append(detectors, d)

	result := make([]*insight.AnalyzerSummary, 0)
	for _, s := range stocks {
		analyzer := insight.NewAnalyzer(client, detectors...)
		indicators := analyzer.Analyze(context.Background(), s)
		if indicators != nil {
			result = append(result, &insight.AnalyzerSummary{s, indicators})
		}
	}
	return result
}
