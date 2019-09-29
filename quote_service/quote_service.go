package main

import (
  "context"
  "log"
  "net"
  "fmt"

  "google.golang.org/grpc"
  pb "stockbuddy/protos/quote_go_proto"
  yahooclient "stockbuddy/quote_service/yahoofinanceclient"
)

type QuoteServer struct {
  client *yahooclient.YahooFinanceClient
}

func NewQuoteServer() *QuoteServer {
return &QuoteServer{
  client: yahooclient.NewYahooFinanceClient(5), // 5s http timeout
}
}

func (q *QuoteServer) ListQuoteHistory(ctx context.Context, req *pb.QuoteRequest) (*pb.QuoteResponse, error) {
  quotes, err := q.client.GetQuoteHistory(req.Symbol, int(req.Period))
  if err != nil {
    log.Println(err.Error())
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
    log.Fatal("Failed to listen on localhost:50051")
  }
  server := grpc.NewServer()
  pb.RegisterQuoteServiceServer(server, NewQuoteServer())
  server.Serve(list)
  fmt.Print("Listening on :50051")
}
