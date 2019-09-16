package main

import (
  "context"
  "errors"
  "fmt"
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
    &quotepb.QuoteRequest{Symbol: "GOOG", Period: 200},
  )

  if err != nil {
    log.Panic(err.Error())
  }

  log.Printf("Num rows returned=%d", len(quoteResponse.Quotes))
  //for _, quote := range quoteResponse.Quotes {
    //log.Printf("close=%v", quote.Close)
  //}
  ma, _ := NDayMovingAverageWithOffset(50, 0, quoteResponse.Quotes)
  log.Print(ma)
  conn.Close()
}

// NDayMovingAverageWithOffset calculates N-day moving average for a quote series ordered
// in ascending sequential order (newest quote last). 
// An offset "X" can be used to calculate the N-day moving average X days ago. Should default
// to 0.
func NDayMovingAverageWithOffset(n int, offset int, series []*quotepb.Quote) (float64, error) {
  seriesLen := len(series)
  if seriesLen < n + offset {
    return 0., errors.New(fmt.Sprintf("Series len must be >= %d, but it is %d", n + offset, seriesLen))
  }
  // Take the N last quotes in the series
  nDaySeries := series[seriesLen-n-offset:seriesLen-offset]
  accum := 0.
  for _, quote := range nDaySeries {
    accum += quote.Close
  }
  return accum/float64(n), nil
}
