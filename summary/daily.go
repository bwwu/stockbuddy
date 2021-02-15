package main

import (
	"context"
	"log"
	"stockbuddy/analysis/detectors/pricedelta"
	"stockbuddy/analysis/insight"
	quotepb "stockbuddy/protos/quote_go_proto"
	"stockbuddy/smtp"
	"time"

	"google.golang.org/grpc"
)

var watchlist []string = []string{"DDD", "PRLB", "SSYS"}

func main() {
	emailPassword := ""
	mailList := []string{""}
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
			log.Printf(`daily: %d indicator(s) found for "%s".`, len(summ.Indicators), summ.Symbol)
		}
		msgBody := insight.FormatByIndicator(summaries)
		log.Printf("daily: analysis took %d ms", time.Now().Sub(t1).Milliseconds())

		mail(emailPassword, msgBody, mailList)
	}
}

func mail(password, content string, recipients []string) {
	subject := "Daily price summary"

	body := "<p>Your watchlist</p>\n" + content

	email := smtp.Email{body, subject, recipients}
	email.Send(password)
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
