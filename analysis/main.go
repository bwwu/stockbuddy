package main

import (
  "context"
  "log"
  "google.golang.org/grpc"
  ma "stockbuddy/analysis/moving_average"
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
  ma, _ := ma.NDayMovingAverageWithOffset(50, 0, quoteResponse.Quotes)
  log.Print(ma)
  conn.Close()
}
