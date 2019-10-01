package main

import (
  "context"
  "log"
  "google.golang.org/grpc"

  "stockbuddy/smtp/sendmail"
  stw "stockbuddy/analysis/stocks_to_watch"
  ma "stockbuddy/analysis/moving_average/moving_average"
  macd "stockbuddy/analysis/detectors/macd_crossover"
  cr "stockbuddy/analysis/moving_average/crossover/crossover_reporter"
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
      crossovers = append(crossovers, summary...)
    }
  }
  if len(crossovers) > 0 {
    subject := "New Crossovers Detected"
    recipients := []string{"brandonwu23@gmail.com"}
    log.Printf("%d crossovers found\n", len(crossovers))

    body := "<p>The following stocks have emitted either a 12-Day/48-Day Simple Moving Average - <strong>SMA(12,48)</strong> or a Moving Average Convergence Divergence - <strong>MACD(12,26,9)</strong> Crossover signal...</p>\n"
    body = body + ma.GetSummaryTable(crossovers)
    email := sendmail.Email{body, subject, recipients}
    email.Send()
  }
}

func calculateMACrossover(c quotepb.QuoteServiceClient, symbol string) []*cr.CrossoverReporter {
  req := &quotepb.QuoteRequest{Symbol: symbol, Period: 365}
  resp, err := c.ListQuoteHistory(context.Background(), req)
  crossovers := make([]*cr.CrossoverReporter, 0)

  if err != nil {
    log.Println(err.Error())
    return []*cr.CrossoverReporter{}
  }

  smaCrossover, err := ma.NewMovingAverageCrossoverReporter(12, 48, resp.Quotes)
  if err != nil {
    log.Println(err.Error())
    return []*cr.CrossoverReporter{}
  }
  if smaCrossover.GetCrossover() != 0 {
    log.Printf("%s MA-Crossover found for \"%s\"", smaCrossover.GetCrossover().String(), symbol)
    crossovers = append(crossovers, smaCrossover)
  }
  macdCrossover, err := macd.DetectMovingAverageConvergenceDivergenceCrossover(resp.Quotes)
  if err != nil {
    log.Println(err.Error())
    return []*cr.CrossoverReporter{}
  }

  // log.Printf("MACD for %s = %v\n", symbol, macdCrossover.SeriesA[len(macdCrossover.SeriesA)-1])
  // log.Printf("Signal line for %s = %v\n", symbol, macdCrossover.SeriesB[len(macdCrossover.SeriesB)-1])
  if macdCrossover.GetCrossover() != 0 {
    log.Printf("%s MACD-Crossover found for \"%s\"", macdCrossover.GetCrossover().String(), symbol)
    crossovers = append(crossovers, macdCrossover)
  }
  return crossovers
}
