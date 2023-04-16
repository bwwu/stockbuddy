package rsi

import (
	"fmt"
	"github.com/bwwu/stockbuddy/analysis/constants"
	"github.com/bwwu/stockbuddy/analysis/lib/sma"
)

/**
 * RelativeStrengthIndexSeries returns N-period RSI given a price series.
 * Price series should ordered in descending order of age (oldest first).
 * Requires that the series length is at least N+2.
 */
func RelativeStrengthIndexSeries(n int, prices []float64) ([]float64, error) {
	// Need a minimum of N+3 prices to calculate N-period RSI.
	if len(prices) < n+2 {
		return nil, fmt.Errorf("rsi: length of price series %d is insuficient for N=%d", len(prices), n)
	}
	// Diff function shifts series +1.
	deltas := diff(prices)
	ups := make([]float64, len(prices)-1)
	downs := make([]float64, len(prices)-1)

	// Daily gains and losses.
	for i, delta := range deltas {
		if delta > 0 {
			ups[i] = delta
		} else if delta < 0 {
			downs[i] = -1. * delta
		}
	}

	// Seed RS w arithmetic mean of first-N ups & downs (first N+1 days).
	avgUp, err := sma.SimpleMovingAverage(n, ups[:n])
	if err != nil {
		return nil, err
	}
	avgDown, err := sma.SimpleMovingAverage(n, downs[:n])
	if err != nil {
		return nil, err
	}

	rsi := make([]float64, len(prices)-n-1)

	// Start at day N+2.
	for i := n; i < len(deltas); i++ {
		// Smooth averages.
		avgUp = (float64(n-1)*avgUp + ups[i]) / float64(n)
		avgDown = (float64(n-1)*avgDown + downs[i]) / float64(n)
		rs := avgUp / avgDown
		rsi[i-n] = 100. - 100./(1.+rs)
	}
	return rsi, nil
}

func RelativeStrengthIndex(n int, prices []float64) (float64, error) {
	series, err := RelativeStrengthIndexSeries(n, prices)
	if err != nil {
		return 0., err
	}
	return series[len(series)-1], nil
}

func ToPriceExtension(rsi float64) constants.PriceExtension {
	if rsi < 30.0 {
		return constants.Oversold
	}
	if rsi > 70.0 {
		return constants.Overbought
	}
	return 0
}

// Calculate price deltas, for t=0,...N, diff(t) = Price(t+1)-Price(t).
func diff(prices []float64) []float64 {
	diffs := make([]float64, len(prices)-1)
	for i := 0; i < len(prices)-1; i++ {
		diffs[i] = prices[i+1] - prices[i]
	}
	return diffs
}
