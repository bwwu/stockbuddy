package main

import (
  "context"
  "log"
  "google.golang.org/grpc"

  "stockbuddy/smtp/sendmail"
  stw "stockbuddy/analysis/stocks_to_watch"
  ma "stockbuddy/analysis/moving_average/moving_average"
  cr "stockbuddy/analysis/moving_average/crossover/crossover_reporter"
  "stockbuddy/analysis/moving_average/crossover/crossover"
  quotepb "stockbuddy/protos/quote_go_proto"
)

func main() {
  conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
  if err != nil {
    log.Fatal(err.Error())
  }
  client := quotepb.NewQuoteServiceClient(conn)
  searchSymbolsForMACrossover(client, stw.StocksToWatch)
  conn.Close()
}

func searchSymbolsForMACrossover(c quotepb.QuoteServiceClient, symbols []string) {
  crossovers := make([]*cr.CrossoverReporter, 0)
  for _, symbol := range symbols {
    summary := calculateMACrossover(c, symbol)
    if summary != nil {
      crossovers = append(crossovers, summary)
    }
  }
  if len(crossovers) > 0 {
    subject := "12/48-Day MA Crossover detected"
    recipients := []string{"brandonwu23@gmail.com"}
    log.Printf("%d crossovers found\n", len(crossovers))

    body := "<p>The following stocks have emitted a 12/48-day MA Crossover signal...</p>\n"
    body = body + ma.GetSummaryTable(crossovers)
    email := sendmail.Email{body, subject, recipients}
    email.Send()
  }
}

func calculateMACrossover(c quotepb.QuoteServiceClient, symbol string) *cr.CrossoverReporter {
  req := &quotepb.QuoteRequest{Symbol: symbol, Period: 365}
  resp, err := c.ListQuoteHistory(context.Background(), req)

  if err != nil {
    log.Println(err.Error())
    return nil
  }

  summary, err := ma.NewMovingAverageCrossoverReporter(12, 48, resp.Quotes)
  if err != nil {
    log.Println(err.Error())
    return nil
  }
  if summary.GetCrossover() == crossover.None {
    return nil
  }
  log.Printf("%s MA-Crossover found for \"%s\"", summary.GetCrossover().String(), symbol)
  return summary
}
