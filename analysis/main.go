package main

import (
  "context"
  "fmt"
  "log"
  "google.golang.org/grpc"

  "stockbuddy/smtp/sendmail"
  ma "stockbuddy/analysis/moving_average"
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
}

func main() {
  conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
  if err != nil {
    log.Fatal(err.Error())
  }
  client := quotepb.NewQuoteServiceClient(conn)

  searchSymbolsForCrossover(client, symbols)

//  quoteResponse, err := client.ListQuoteHistory(
//    context.Background(),
//    &quotepb.QuoteRequest{Symbol: "GOOG", Period: 365},
//  )
//
//  log.Printf("Num rows returned=%d", len(quoteResponse.Quotes))
//  ma, _ := ma.NDayMovingAverageWithOffset(50, 0, quoteResponse.Quotes)
//  log.Print(ma)
  conn.Close()
}

var heading = "<tr><th>SYM</th><th>12DMA</th><th>12DMAΔ</th><th>48DMA</th><th>48DMAΔ</th><th>SIGNAL</th></tr>\n"

func searchSymbolsForCrossover(c quotepb.QuoteServiceClient, symbols []string) {
  crossovers := make([]*ma.MovingAverageCrossoverSummary, 0)
  for _, symbol := range symbols {
    summary := calculateMACrossoverForSymbol(c, symbol)
    if summary != nil {
      crossovers = append(crossovers, summary)
    }
  }
  if len(crossovers) > 0 {
    subject := "12/48-Day MA Crossover detected"
    recipients := []string{"brandonwu23@gmail.com"}
    log.Printf("%d crossovers found\n", len(crossovers))
    body := "<p>The following symbols have emitted a 12/48-day MA Crossover signal</p>"
    body = body + "<table  cellspacing=\"0\" cellpadding=\"0\" width=\"640\" align=\"center\" border=\"1\">\n" + heading
    for _, c := range crossovers {
      body = body + formatMACrossoverRow(c)
    }
    body = body + "</table>"
    email := sendmail.Email{body, subject, recipients}
    email.Send()
  }
}

func formatMACrossoverRow(s *ma.MovingAverageCrossoverSummary) string {
  shortDelta := s.ShortMA - s.ShortMAMinus1
  longDelta := s.LongMA - s.LongMAMinus1

  var signal string
  if s.Crossover == ma.Bullish {
    signal = "BUY"
  } else {
    signal = "SELL"
  }

  template := "<tr><td>%s</td><td>%.2f</td><td>%.2f</td><td>%.2f</td><td>%.2f</td><td>%s</td></tr>\n"
  return fmt.Sprintf(template, s.Symbol, s.ShortMA, shortDelta, s.LongMA, longDelta, signal)
}

func calculateMACrossoverForSymbol(c quotepb.QuoteServiceClient, symbol string) *ma.MovingAverageCrossoverSummary {
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
