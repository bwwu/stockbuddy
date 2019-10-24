package yahoofinance

import (
  "fmt"
  "strconv"
  "context"
  "time"
  "encoding/csv"
  "net/http"
)

type YahooFinanceClient struct {
  client http.Client
}

type QuoteHistory struct {
  Symbol string
  DailyQuotes []DailyQuote
}

type DailyQuote struct {
  Timestamp time.Time
  Open, High, Low, Close, AdjClose float64
  Volume uint32
}

func NewYahooFinanceClient(timeoutInSec int) *YahooFinanceClient {
  return &YahooFinanceClient{
    client: http.Client{
      Timeout: time.Duration(timeoutInSec) * time.Second,
    },
  }
}

func (y *YahooFinanceClient) GetQuoteHistory(ctx context.Context, symbol string, days int) (*QuoteHistory, error) {
  cookies := map[string]string{
    "B": "ajch0f5elj4sp",
    "APID": "1Adf2ce59c-c1e2-11e9-adce-025f25c4bfdc",
    "APIDTS": "1566870258",
    "PRF": "t%3D" + symbol,
  }
  timeEnd := time.Now()
  timeStart := timeEnd.AddDate(0,0,-1*days)

  // Params: symbol, start timestamp, end timestamp
  const urlFormat = "https://query1.finance.yahoo.com/v7/finance/download/%s?period1=%d&period2=%d&interval=1d&events=history&crumb=pRB6UiIiFnn"

  req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf(urlFormat, symbol, timeStart.Unix(), timeEnd.Unix()), nil)
  if err != nil {
    return nil, err
  }

  for key,val := range cookies {
    req.AddCookie(&http.Cookie{
      Name: key,
      Value: val,
    })
  }

  resp, err := y.client.Do(req)
  if err != nil {
    return nil, err
  }

  defer resp.Body.Close()
  reader := csv.NewReader(resp.Body)
  rows, err := reader.ReadAll()
  if err != nil {
    return nil, err
  }

  rows = rows[1:]
  history := QuoteHistory{
    Symbol: symbol,
    DailyQuotes: make([]DailyQuote, len(rows)),
  }

  for i, row := range rows {
    // Values open, high, low, close, adj_close should be floats
    // TODO don't ignore errors
    floats, _ := parseFloats(row[1:6])
    volume, _ := strconv.ParseUint(row[6], 10, 32)
    timestamp, _ := time.Parse("2006-01-02", row[0])
    history.DailyQuotes[i] = DailyQuote{
      Timestamp: timestamp,
      Open: floats[0],
      High: floats[1],
      Low: floats[2],
      Close: floats[3],
      AdjClose: floats[4],
      Volume: uint32(volume),
    }
  }
  return &history, nil
}

func parseFloats(strs []string) ([]float64, error) {
  floats := make([]float64, len(strs))
  for i, str := range(strs) {
    float, err := strconv.ParseFloat(str, 64)
    if err != nil {
      return nil, err
    }
    floats[i] = float
  }
  return floats, nil
}
