package main

import (
  "context"
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
    log.Printf("close=%f", quote.Close)
  } 
  
  conn.Close()
}
