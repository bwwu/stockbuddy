package main

import (
  "context"
  "log"
  "net"

  "google.golang.org/grpc"
  pb "stockbuddy/protos/quote_go_proto"
  "stockbuddy/quote_service/yahoofinance"
)

type QuoteServer struct {
  client *yahoofinance.YahooFinanceClient
}

func NewQuoteServer() *QuoteServer {
  return &QuoteServer{
    client: yahoofinance.NewYahooFinanceClient(5), // 5s http timeout
  }
}

func (q *QuoteServer) ListQuoteHistory(ctx context.Context, req *pb.QuoteRequest) (*pb.QuoteResponse, error) {
  quotes, err := q.client.GetQuoteHistory(req.Symbol, int(req.Period))
  if err != nil {
    log.Printf("quote_service: ListQuoteHistory(%v,%v) error\n,%v", req.Symbol, req.Period, err)
    return nil, err
  }

  quoteProtos := make([]*pb.Quote, len(quotes.DailyQuotes))
  for i, quote := range quotes.DailyQuotes {
    quoteProtos[i] = &pb.Quote{
      Symbol: req.Symbol,
      Timestamp: quote.Timestamp.Unix(),
      Open: quote.Open,
      High: quote.High,
      Low: quote.Low,
      Close: quote.Close,
      AdjClose: quote.AdjClose,
      Volume: quote.Volume,
    }
  }

  return &pb.QuoteResponse{Quotes: quoteProtos}, nil
}

func main() {
  list, err := net.Listen("tcp", "localhost:50051")
  if err != nil {
    log.Fatalf("Failed to listen on :50051 due to error: %s", err.Error())
  }
  server := grpc.NewServer()
  pb.RegisterQuoteServiceServer(server, NewQuoteServer())

  log.Println("Starting quote service on localhost:50051...")
  if err := server.Serve(list); err != nil {
    log.Fatal(err.Error())
  }
}
