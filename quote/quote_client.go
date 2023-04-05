package quote

import (
	"context"
	"log"
	"github.com/bwwu/stockbuddy/quote/yahoofinance"
)

type QuoteClient struct {
	client *yahoofinance.YahooFinanceClient
}

func NewQuoteClient() QuoteClient {
	return QuoteClient{
		client: yahoofinance.NewYahooFinanceClient(5), // 5s http timeout
	}
}

func (q *QuoteClient) ListQuoteHistory(ctx context.Context, symbol string, period int) ([]*Quote, error) {
	yfquotes, err := q.client.GetQuoteHistory(ctx, symbol, period)
	if err != nil {
		log.Printf("ListQuoteHistory(%v,%v) error\n,%v", symbol, period, err)
		return nil, err
	}

	quotes := make([]*Quote, len(yfquotes))
	for i, q := range yfquotes {
		quotes[i] = &Quote{
			Symbol:    symbol,
			Timestamp: q.Timestamp.Unix(),
			Open:      q.Open,
			High:      q.High,
			Low:       q.Low,
			Close:     q.Close,
			AdjClose:  q.AdjClose,
			Volume:    q.Volume,
		}
	}

	return quotes, nil
}

