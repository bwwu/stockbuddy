package main

import (
  "context"
  "fmt"
  "google.golang.org/grpc"
  "log"

  "stockbuddy/analysis/insight"
  detectors "stockbuddy/analysis/detectors/sma_crossover"
  pb "stockbuddy/protos/quote_go_proto"
)

func main() {
  conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
  if err != nil {
    log.Fatal(err.Error())
  }
  client := pb.NewQuoteServiceClient(conn)

  detector,_ := detectors.NewSimpleMovingAverageDetector(12, 48)
  analyzer := insight.NewAnalyzer(client, detector)
  log.Print(detector)

  summary := analyzer.Analyze(context.Background(),"JNJ")
  fmt.Println(summary[0].Outlook())
  conn.Close()
}
