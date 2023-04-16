// Package macdrsi implements a compound detector combining signals from a MACD
// crossover and RSI.
package macdrsi

import (
	"fmt"
	"github.com/bwwu/stockbuddy/analysis/constants"
	"github.com/bwwu/stockbuddy/analysis/detectors/macdx"
	"github.com/bwwu/stockbuddy/analysis/lib/rsi"
	"github.com/bwwu/stockbuddy/analysis/insight"
	"github.com/bwwu/stockbuddy/quote"
	"github.com/bwwu/stockbuddy/util"


)

// A MACDxRSIDetector is a compound detector combining MACD crossover and RSI
//	swing rejection.
type MACDxRSIDetector struct
{
	period, lookback int // Num days over which to smooth RSI.
	macdx *macdx.MACDDetector
}

func NewMACDxRSIDetector(shortTerm, longTerm, signalTerm, period, lookback int) (*MACDxRSIDetector, error) {
	macdx, err := macdx.NewMACDDetector(shortTerm, longTerm, signalTerm)
	if err != nil {
		return nil, fmt.Errorf("macd_rsi: invalid params for macd: '%v'", err)
	}

	return &MACDxRSIDetector{period, lookback, macdx}, nil
}

func (d *MACDxRSIDetector) Process(quotes []*quote.Quote) (insight.Indicator, error) {
	// TODO handle errs
	macdInsight, _ := d.macdx.Process(quotes)
	extensions, _ := d.generateRSIPriceExtension(quotes)

	if macdInsight == nil {
		return nil, nil
	}

	for i:= 0; i < d.lookback; i++ {
		ext := extensions[len(extensions)-i-1]
		outlook := macdInsight.Outlook()
		if ext == constants.Oversold && outlook == constants.Bullish || ext == constants.Overbought && outlook == constants.Bearish {
			return &MACDxRSIIndicator{
				outlook,
			}, nil
		}
	}

	return nil, nil
}

func (d *MACDxRSIDetector) generateRSIPriceExtension(quotes []*quote.Quote) ([]constants.PriceExtension, error) {
	prices := make([]float64, 0, len(quotes))
	for _, q := range quotes {
		prices = append(prices, q.Close)
	}
	rsiSeries, err := rsi.RelativeStrengthIndexSeries(d.period, prices)
	if err != nil {
		return nil, err
	}
	return util.Map(rsi.ToPriceExtension, rsiSeries), nil
}

type MACDxRSIIndicator struct {
	outlook constants.Outlook
}

func (i MACDxRSIIndicator) Name() string {
	return fmt.Sprintf("MACD-RSI composite")
}

func (i MACDxRSIIndicator) Summary() string {
	return fmt.Sprintf("")
}

func (i MACDxRSIIndicator) Outlook() constants.Outlook {
	return i.outlook
}

func (i MACDxRSIIndicator) Trend() constants.Trend {
	return constants.Reversal
}