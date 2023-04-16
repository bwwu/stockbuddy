package macdx

import (
	"fmt"
	"github.com/bwwu/stockbuddy/analysis/constants"
	"github.com/bwwu/stockbuddy/analysis/insight"
	"github.com/bwwu/stockbuddy/analysis/lib/crossover"
	"github.com/bwwu/stockbuddy/analysis/lib/macd"
	"github.com/bwwu/stockbuddy/analysis/lib/sma"
	"github.com/bwwu/stockbuddy/quote"
)

// MACDDetector
type MACDDetector struct {
	shortTerm, longTerm, signalTerm int
}

func NewMACDDetector(shortTerm, longTerm, signalTerm int) (*MACDDetector, error) {
	if shortTerm >= longTerm || shortTerm <= 0 {
		return nil, fmt.Errorf(
			"macd_crossover: invalid long(%d) and short(%d) term values for MACD",
			longTerm,
			shortTerm,
		)
	}
	return &MACDDetector{
		shortTerm:  shortTerm,
		longTerm:   longTerm,
		signalTerm: signalTerm,
	}, nil
}

func (d *MACDDetector) Process(quotes []*quote.Quote) (insight.Indicator, error) {
	prices := make([]float64, 0, len(quotes))
	for _, q := range quotes {
		prices = append(prices, q.Close)
	}

	crossovers, err := d.GenerateMACDCrossoverSeries(prices)
	if err != nil {
		return nil, err
	}
	
	recentCrossover := crossovers[len(crossovers)-1]
	if recentCrossover == 0 {
		return nil, nil
	}
	return &MACDCrossover{
		shortTerm:  d.shortTerm,
		longTerm:   d.longTerm,
		signalTerm: d.signalTerm,
		outlook:    recentCrossover,
	}, nil
}

// GenerateMACDCrossoverSeries returns a list of MACD/Signal line crossover
//	events. For example, when a MACD 12-26 line positively crosses a 9 day
//	signal line, emit a BULLISH outlook. The outlook at the len(series)-1
//	index represents the most recent day's outlook.
func (d *MACDDetector) GenerateMACDCrossoverSeries(prices []float64) ([]constants.Outlook, error) {
	macdSeries, err := macd.MovingAverageConvergenceDivergenceSeries(d.shortTerm, d.longTerm, prices)
	if err != nil {
		return nil, err
	}
	signalLine, err := sma.ExponentialMovingAverageSeries(d.signalTerm, macdSeries)
	if err != nil {
		return nil, err
	}

	return crossover.DetectCrossovers(macdSeries, signalLine), nil
}

// MACDCrossover
type MACDCrossover struct {
	shortTerm, longTerm, signalTerm int
	outlook                         constants.Outlook
}

func (c MACDCrossover) Name() string {
	return fmt.Sprintf("MACD(%d,%d,%d)", c.shortTerm, c.longTerm, c.signalTerm)
}

func (c MACDCrossover) Summary() string {
	return fmt.Sprintf(
		"MACD(%d,%d,%d)",
		c.shortTerm,
		c.longTerm,
		c.signalTerm,
	)
}

func (c MACDCrossover) Outlook() constants.Outlook {
	return c.outlook
}

func (c MACDCrossover) Trend() constants.Trend {
	return constants.Reversal
}
