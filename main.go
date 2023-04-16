package main

import (
	"context"
	"flag"
	"log"
	"os"
	"time"

	"github.com/bwwu/stockbuddy/analysis/detectors"
	"github.com/bwwu/stockbuddy/analysis/insight"
	"github.com/bwwu/stockbuddy/fileio"
	"github.com/bwwu/stockbuddy/quote"
	"github.com/bwwu/stockbuddy/smtp"
)

var (
	flagNomail    = flag.Bool("nomail", false, "Set to true to disable sending email reports.")
	flagWatchlist = flag.String("use_watchlist", "watchlists/default.txt", "Path to txt file with a list of stocks to track.")
	flagMailList  = flag.String("mail_to", "", "Comma separated list of email addresses to whom results will be mailed.")
)

func main() {
	flag.Parse()

	watchlist, err := fileio.ReadLines(*flagWatchlist)
	if err != nil {
		log.Fatal(err.Error())
	}

	var mailList []string
 
	if !*flagNomail {
		mailList, err = smtp.ParseEmailsFromList(*flagMailList)
		if err != nil {
			log.Fatal(err)
		}
	}

	emailPassword := os.Getenv("STOCKBUDDY_PASSWORD")
	if emailPassword == "" && !*flagNomail {
		log.Fatal("main: password unset. Set email credentials on env $STOCKBUDDY_PASSWORD.")
	}

	t1 := time.Now()
	client := quote.NewQuoteClient()

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

func process(client quote.QuoteClient, stocks []string) []*insight.AnalyzerSummary {
	// Instantiate all of the detectors to run.
	detecs, err := detectors.GetDefaultDetectors([]string{
		"macd_rsi",
	}) 
	if err != nil {
		log.Fatal(err)
	} 

	// Spawn goroutine to run analyzer over all detectors, one per stock.
	summaryc := make(chan *insight.AnalyzerSummary)
	defer close(summaryc)

	for _, symbol := range stocks {
		go func(s string) {
			analyzer := insight.NewAnalyzer(client, detecs...)
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
