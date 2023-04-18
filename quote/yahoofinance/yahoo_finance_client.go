package yahoofinance

import (
	"context"
	"encoding/csv"
	"fmt"
	"net/http"
	"strconv"
	"time"
	"golang.org/x/sync/semaphore"
)

type YahooFinanceClient struct {
	client http.Client
	workerPool *semaphore.Weighted
}

type YFQuote struct {
	Timestamp                        time.Time
	Open, High, Low, Close, AdjClose float64
	Volume                           uint64
}

func NewYahooFinanceClient(timeoutInSec, maxWorkers int) *YahooFinanceClient {
	return &YahooFinanceClient{
		client: http.Client{
			Timeout: time.Duration(timeoutInSec)*time.Second,
			Transport: &http.Transport{
				ResponseHeaderTimeout: time.Duration(timeoutInSec)*time.Second,
			},
		},
		workerPool: semaphore.NewWeighted(int64(maxWorkers)),
	}
}

func (y *YahooFinanceClient) GetQuoteHistory(ctx context.Context, symbol string, days int) ([]*YFQuote, error) {
	if err := y.workerPool.Acquire(ctx, 1); err != nil {
		return nil, fmt.Errorf(
			"yahoo_finance_client::GetQuoteHistory: failed to aquire semaphore: %v",
			err,
		)
	}
	defer y.workerPool.Release(1)

	timeEnd := time.Now()
	timeStart := timeEnd.AddDate(0, 0, -1*days)

	// Params: symbol, start timestamp, end timestamp
	const urlFormat = "https://query1.finance.yahoo.com/v7/finance/download/%s?period1=%d&period2=%d&interval=1d&events=history&crumb=pRB6UiIiFnn"
	reqUrl := fmt.Sprintf(urlFormat, symbol, timeStart.Unix(), timeEnd.Unix())
	req, err := http.NewRequestWithContext(ctx, "GET", reqUrl, nil)
	if err != nil {
		return nil, err
	}

	resp, err := y.client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.Status != "200 OK" {
		return nil, fmt.Errorf(
			`yahoo_finance_client::GetQuoteHistory: requested failed with status=%v, url="%v"`,
			resp.Status,
			reqUrl,
		)
	}

	defer resp.Body.Close()
	reader := csv.NewReader(resp.Body)
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	rows = rows[1:]
	quotes := make([]*YFQuote, len(rows))

	for i, row := range rows {
		// Values open, high, low, close, adj_close should be floats
		// TODO don't ignore errors
		floats, err := parseFloats(row[1:6])
		if err != nil {
			return nil, err
		}

		volume, err := strconv.ParseUint(row[6], 10, 64)
		if err != nil {
			return nil, err
		}

		timestamp, err := time.Parse("2006-01-02", row[0])
		if err != nil {
			return nil, err
		}

		quotes[i] = &YFQuote{
			Timestamp: timestamp,
			Open:      floats[0],
			High:      floats[1],
			Low:       floats[2],
			Close:     floats[3],
			AdjClose:  floats[4],
			Volume:    uint64(volume),
		}
	}
	return quotes, nil
}

func parseFloats(strs []string) ([]float64, error) {
	floats := make([]float64, len(strs))
	for i, str := range strs {
		float, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return nil, err
		}
		floats[i] = float
	}
	return floats, nil
}
