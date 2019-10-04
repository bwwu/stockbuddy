package main

import (
  "context"
  "log"
  "google.golang.org/grpc"

  "stockbuddy/smtp/sendmail"
  "stockbuddy/analysis/insight"
  sma "stockbuddy/analysis/detectors/sma_crossover"
  macd "stockbuddy/analysis/detectors/macd_crossover"
  pb "stockbuddy/protos/quote_go_proto"
)

func main() {
  conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
  if err != nil {
    log.Fatal(err.Error())
  }
  client := pb.NewQuoteServiceClient(conn)
  summaries := process(client, StocksToWatch)
  if len(summaries) > 0 {
    for _, summ := range summaries {
      log.Printf("main: %d indicators found for %s.\n", len(summ.Indicators), summ.Symbol)
    }
    msgBody := insight.TableFormat(summaries)
    mail(msgBody)
  }
  conn.Close()
}

func mail(content string) {
  subject := "Reversal Trends Detected"
  recipients := []string{"brandonwu23@gmail.com"}

  body := "<p>Reversal trends have been detected for the following stocks:</p>\n" + content

  email := sendmail.Email{body, subject, recipients}
  email.Send()
}

func process(client pb.QuoteServiceClient, stocks []string) []*insight.AnalyzerSummary {
  // Instantiate all of the detectors to run.
  detectors := make([]insight.Detector, 0)
  if smaDetec, err := sma.NewSimpleMovingAverageDetector(12, 48); err != nil {
    log.Fatal(err)
  } else {
    detectors = append(detectors, smaDetec)
  }
  if macdDetec, err := macd.NewMACDDetector(12, 26, 9); err != nil {
    log.Fatal(err)
  } else {
    detectors = append(detectors, macdDetec)
  }

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
  for i:=0; i<len(stocks); i++ {
    summary := <-summaryc
    if summary != nil {
      result = append(result, summary)
    }
  }
  return result
}
