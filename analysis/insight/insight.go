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
	Process([]*pb.Quote) (Indicator, error)
}

type AnalyzerSummary struct {
	Symbol     string
	Indicators []Indicator
}

// Analyzer represents a series of computation which the caller can inoke
// across multiple symbols.
type Analyzer struct {
	client    pb.QuoteServiceClient
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
		return nil
	}

	indicatorc := make(chan Indicator)
	errc := make(chan error)
	defer close(indicatorc)
	defer close(errc)

	// Spawn goroutine per detector
	for _, d := range a.detectors {
		go func(detector Detector) {
			if ind, err := detector.Process(resp.Quotes); err != nil {
				errc <- err
			} else {
				indicatorc <- ind
			}
		}(d)
	}

	// Collect errors and results from indicators
	indicators := make([]Indicator, 0)
	for i := 0; i < len(a.detectors); i++ {
		select {
		case indic := <-indicatorc:
			if indic != nil {
				indicators = append(indicators, indic)
			}
		case err := <-errc:
			log.Println(err.Error())
		}
	}
	if len(indicators) == 0 {
		return nil
	}
	return indicators
}
