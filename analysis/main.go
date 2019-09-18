package main

import (
  "context"
  "log"
  "google.golang.org/grpc"

  "stockbuddy/smtp/sendmail"
  ma "stockbuddy/analysis/moving_average/moving_average"
  quotepb "stockbuddy/protos/quote_go_proto"
)

var symbols = []string{
  "AXP",
  "AAPL",
  "BA",
  "CAT",
  "CSCO",
  "CVX",
  "XOM",
  "GS",
  "HD",
  "IBM",
  "INTC",
  "JNJ",
  "KO",
  "JPM",
  "MCD",
  "MMM",
  "MRK",
  "MSFT",
  "NKE",
  "PFE",
  "PG",
  "TRV",
  "UNH",
  "UTX",
  "VZ",
  "V",
  "WBA",
  "WMT",
  "DIS",
  "DOW",
  "WM",
  "HQY",
  "SNPS",
  "WIX",
  "ZNGA",
  "OLLI",
  "HRC",
  "TWLO",
  "NTDOY",
  "APPN",
  "HA",
  "TWLO",
  "TLK",
  "MA",
  "ZBRA",
  "ZS",
  "AMZN",
  "SQ",
  "UNP",
  "SFIX",
  "ALK",
  "ANET",
  "BB",
  "SFIX",
  "AMGN",
  "WIX",
  "NEWR",
  "SHOP",
  "SHOP",
  "OKTA",
  "MKL",
  "CRUS",
  "VRNS",
  "FICO",
  "PAYC",
  "OKTA",
}

func main() {
  conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
  if err != nil {
    log.Fatal(err.Error())
  }
  client := quotepb.NewQuoteServiceClient(conn)
  searchSymbolsForMACrossover(client, symbols)
  conn.Close()
}

func searchSymbolsForMACrossover(c quotepb.QuoteServiceClient, symbols []string) {
  crossovers := make([]*ma.MovingAverageCrossoverSummary, 0)
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

func calculateMACrossover(c quotepb.QuoteServiceClient, symbol string) *ma.MovingAverageCrossoverSummary {
  req := &quotepb.QuoteRequest{Symbol: symbol, Period: 365}
  resp, err := c.ListQuoteHistory(context.Background(), req)

  if err != nil {
    log.Println(err.Error())
    return nil
  }

  summary, err := ma.NewMovingAverageCrossoverSummary(12, 48, resp.Quotes)
  if err != nil {
    log.Println(err.Error())
    return nil
  }
  if summary.Crossover == ma.None {
    return nil
  }
  log.Printf("%s MA-Crossover found for \"%s\"", summary.Crossover.String(), symbol)
  return summary
}
