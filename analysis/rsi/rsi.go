package rsi

import (
    "fmt"
    "errors"
)

/**
 * RelativeStrengthIndex calculates N-period RSI for a price series.
 * Price series should ordered in descending order of age (oldest first).
 * Requires that the series length is at least N+2.
 */
func RelativeStrengthIndex(n int, prices []float64) (float64, error) {
  // Need a minimum of N+3 prices to calculate N-period RSI.
  if len(prices) < n+2 {
    return 0., errors.New(fmt.Sprintf("length of price series %d is insuficient for N=%d", len(prices), n))
  }
  // diff function shifts series +1
  deltas := diff(prices)
  ups := make([]float64, len(prices)-1)
  downs := make([]float64, len(prices)-1)

  // Daily gains and losses.
  for i, delta := range deltas {
    if (delta > 0) {
      ups[i] = delta
    } else if (delta < 0) {
      downs[i] = -1.*delta
    }
  }

  // Seed RS with average of first-N ups & downs (first N+1 days)
  avgUp := average(ups[:n])
  avgDown := average(downs[:n])

  // Starting at day N+2
  for i := n; i < len(deltas); i++ {
    // Calculate "smoothed" averages
    avgUp = (float64(n-1)*avgUp + ups[i])/float64(n)
    avgDown = (float64(n-1)*avgDown + downs[i])/float64(n)
  }
  rs := avgUp/avgDown
  return 100. - 100./(1.+rs),nil
}

func average(values []float64) float64 {
  accum := 0.
  for _, val := range values {
    accum += val
  }
  return accum/float64(len(values))
}

func diff(prices []float64) []float64 {
  diffs := make([]float64, len(prices)-1)
  for i:=0; i<len(prices)-1; i++ {
    diffs[i] = prices[i+1] - prices[i]
  }
  return diffs
}
