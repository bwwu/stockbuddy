package main

import (
  "context"
  "log"
  "net"
  "google.golang.org/grpc"
  "stockbuddy/protos/quote"
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

func (q *QuoteServer) ListQuoteHistory(ctx context.Context, req *quote.QuoteRequest) (*quote.QuoteResponse, error) {
  quotes, err := q.client.GetQuoteHistory(req.Symbol, int(req.Period))
  if err != nil {
    log.Printf("quote_service: ListQuoteHistory(%v,%v) error\n,%v", req.Symbol, req.Period, err)
    return nil, err
  }

  quoteProtos := make([]*quote.Quote, len(quotes.DailyQuotes))
  for i, q := range quotes.DailyQuotes {
    quoteProtos[i] = &quote.Quote{
      Symbol: req.Symbol,
      Timestamp: q.Timestamp.Unix(),
      Open: q.Open,
      High: q.High,
      Low: q.Low,
      Close: q.Close,
      AdjClose: q.AdjClose,
      Volume: q.Volume,
    }
  }

  return &quote.QuoteResponse{Quotes: quoteProtos}, nil
}

func main() {
  list, err := net.Listen("tcp", "localhost:50051")
  if err != nil {
    log.Fatalf("Failed to listen on :50051 due to error: %s", err.Error())
  }
  server := grpc.NewServer()
  quote.RegisterQuoteServiceServer(server, NewQuoteServer())

  log.Println("Starting quote service on localhost:50051...")
  if err := server.Serve(list); err != nil {
    log.Fatal(err.Error())
  }
}
