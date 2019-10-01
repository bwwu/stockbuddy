// Package insight abstracts the result of an analyzer
package insight

import (
  "context"
  "log"

  "stockbuddy/analysis/constants"
  pb "stockbuddy/protos/quote_go_proto"
)

type Indicator interface {
  // Name of the indicator
  Name() string
  // Relevant values summarizing the result
  Summary() string
  // Whether the result is Bearish or Bullish
  Outlook() constants.Outlook
  // Either Reversal Continuation trend
  Trend() constants.Trend
}

// Detector is a generic type. It computes the val of an indicator and provides
// an interpretation, if any (e.g. Bearish reversal)
type Detector interface {
  Process([]*pb.Quote) (bool, error)
  Get() Indicator
}

// Analyzer represents a series of computation which the caller can inoke
// across multiple symbols.
type Analyzer struct {
  client pb.QuoteServiceClient
  detectors []Detector
}

func NewAnalyzer(client pb.QuoteServiceClient, detectors ...Detector) *Analyzer {
  return &Analyzer{client, detectors}
}

func (a *Analyzer) Analyze(ctx context.Context, symbol string) []Indicator {
  req := &pb.QuoteRequest{Symbol: symbol, Period: 365}
  resp, err := a.client.ListQuoteHistory(ctx, req)
  if err != nil {
    log.Println(err.Error())
    return []Indicator{}
  }

  indicatorc := make(chan Indicator)
  errc := make(chan error)
  defer close(indicatorc)
  defer close(errc)

  // Spawn goroutine per detector
  for _, d := range a.detectors {
    go func() {
      if has, err := d.Process(resp.Quotes); err != nil {
        errc <- err
      } else {
        indicatorc <- d.Get()
      }
    }()
  }

  // Collect errors and results from indicators
  indicators := make([]Indicator, 0)
  for i:=0; i<len(a.detectors); i++ {
    select {
    case indic := <-indicatorc:
      indicators = append(indicators, indic)
    case err := <-errc:
      log.Println(err.Error())
    }
  }
  return indicators
}
