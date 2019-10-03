package main

import (
  "context"
  "google.golang.org/grpc"
  "log"
  "stockbuddy/analysis/insight"
  stw "stockbuddy/analysis/stocks_to_watch"
  sma "stockbuddy/analysis/detectors/sma_crossover"
  macd "stockbuddy/analysis/detectors/macd_crossover2"
  pb "stockbuddy/protos/quote_go_proto"
)

type summary struct {
  symbol string
  indicators []insight.Indicator
}

func main() {
  conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
  if err != nil {
    log.Fatal(err.Error())
  }
  client := pb.NewQuoteServiceClient(conn)
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

  summaryc := make(chan *summary)
  defer close(summaryc)

  for _, symbol := range stw.StocksToWatch {
    go func(s string) {
      analyzer := insight.NewAnalyzer(client, detectors...)
      indicators := analyzer.Analyze(context.Background(), s)
      if indicators == nil {
        summaryc <- nil
      } else {
        summaryc <- &summary{s, indicators}
      }
    }(symbol)
  }

  for i:=0; i<len(stw.StocksToWatch); i++ {
    summary := <-summaryc
    if summary != nil {
      for _, ind := range summary.indicators {
        log.Printf(
          "Found for %s: %s, %s",
          summary.symbol,
          ind.Summary(),
          ind.Outlook().String(),
        )
      }
    }
  }

   //log.Print(detector)

  //summary := analyzer.Analyze(context.Background(),"JNJ")
  //fmt.Println(summary[0].Outlook())
  conn.Close()
}
