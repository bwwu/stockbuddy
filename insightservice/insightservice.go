package main

import (
  "context"
  "log"
  "net"
  "google.golang.org/grpc"
  pb "stockbuddy/protos/insight_go_proto"
)

type InsightService struct {}

func NewInsightService() *InsightService {
  return &InsightService{}
}

func (is *InsightService) AddInsight(ctx context.Context, req *pb.AddInsightRequest) (*pb.AddInsightResponse, error) {
  return nil, nil
}

func (is *InsightService) FetchInsight(ctx context.Context, req *pb.FetchInsightRequest) (*pb.FetchInsightResponse, error) {
  return nil, nil
}

func main() {
  rpcAddr := "localhost:50052"
  list, err := net.Listen("tcp", rpcAddr)
  if err != nil {
    log.Fatalf("insightservice: net.Listen(...) failed with error: %v", err)
  }
  server := grpc.NewServer()
  pb.RegisterInsightStoreServer(server, NewInsightService())

  log.Printf("insightservice: starting on %v\n", rpcAddr)
  if err := server.Serve(list); err != nil {
    log.Fatalf("insightservice: server.Serve(...) failed with error: %v", err)
  }
}
