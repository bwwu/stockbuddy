package pricedelta

import (
	"fmt"
	"stockbuddy/analysis/constants"
	"stockbuddy/analysis/insight"
	pb "stockbuddy/protos/quote_go_proto"
)

// Detector is a summary detector, which always emits an indicator with the price delta over a period.
type Detector struct {
	period int // Num days to look back (default 0)
}

// NewDetector creates a PriceDelta detector.
func NewDetector(period int) (*Detector, error) {
	// Can't look back negative days. Also restrict to 1yr.
	if period < 1 || period > 253 {
		return nil, fmt.Errorf(`pricedelta: cannot period "%d" days`, period)
	}
	return &Detector{period}, nil
}

// NewDefaultDetector creates a price delta detector over today (close - open)
func NewDefaultDetector() *Detector {
	d, _ := NewDetector(1)
	return d
}

// Process subtracts open price from "period" trading days ago from the close price
// of the most recent trading day.
func (d *Detector) Process(quotes []*pb.Quote) (insight.Indicator, error) {
	if len(quotes) < d.period+1 {
		return nil, fmt.Errorf(
			"pricedelta: not enough quotes (%d) in series to compute %d-day delta",
			len(quotes),
			d.period,
		)
	}
	close := quotes[len(quotes)-1].Close
	prev := quotes[len(quotes)-d.period-1].Close

	outlook := constants.Bearish
	if close > prev {
		outlook = constants.Bullish
	}
	return &PriceDelta{
		d.period,
		close - prev,
		(close - prev) * 100. / prev,
		outlook,
	}, nil
}

// PriceDelta is the delta in price from market close on the most recent trading day and
// the opening price from the trading day at the beginning of the period.
type PriceDelta struct {
	period  int
	delta   float64
	percent float64
	outlook constants.Outlook
}

// Name returns "<x>D (+/-)".
func (p PriceDelta) Name() string {
	return fmt.Sprintf("%d-Day (+/-)", p.period)
}

// Summary returns magnitude of price delta and % change.
func (p PriceDelta) Summary() string {
	if p.delta > 0. {
		return fmt.Sprintf("+%.2f (+%.2f%%)", p.delta, p.percent)
	}
	return fmt.Sprintf("%.2f (%.2f%%)", p.delta, p.percent)
}

// Outlook returns "Bullish" for positive delta.
func (p PriceDelta) Outlook() constants.Outlook {
	return p.outlook
}

// Trend returns "Neither".
func (p PriceDelta) Trend() constants.Trend {
	return constants.Neither
}
