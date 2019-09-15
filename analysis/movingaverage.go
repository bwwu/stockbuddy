package main

import (
  "context"
  "errors"
  "log"
  "google.golang.org/grpc"
  quotepb "stockbuddy/protos/quote_go_proto"
)

func main() {
  conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
  if err != nil {
    log.Fatal(err.Error())
  }
  client := quotepb.NewQuoteServiceClient(conn)

  quoteResponse, err := client.ListQuoteHistory(
    context.Background(),
    &quotepb.QuoteRequest{Symbol: "GOOG", Period: 365},
  )

  if err != nil {
    log.Panic(err.Error())
  }

  log.Printf("Num rows returned=%d", len(quoteResponse.Quotes))
  for _, quote := range quoteResponse.Quotes {
    log.Printf("close=%v", quote.Close)
  }
  conn.Close()
}

func NDayMovingAverage(n int, series []quotepb.Quote) (float64, error) {
  if len(series) < n {
    return 0., errors.New("Series len must be >= n")
  }
  accum := float64(0.0)
  for _, quote := range series {
    accum += quote.Close
  }

  return accum/float64(len(series)), nil
}
