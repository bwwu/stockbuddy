// Package macd implements Moving-Average Convergence Divergence alg on a price
// series ordered in descending order of age.
package macd

import (
  "fmt"
  "github.com/bwwu/stockbuddy/analysis/lib/sma"
)

// MovingAverageConvergenceDivergenceSeries returns the difference between
// 12-Period and 26-Period EMA. Generally compared against a Signal line 
// generated from a 9-Day EMA of the MACD.
func MovingAverageConvergenceDivergenceSeries(shortTerm, longTerm int, prices []float64) ([]float64, error) {
  if shortTerm >= longTerm || shortTerm <= 0 {
    return nil, fmt.Errorf("macd: invalid long(%d) and short(%d) terms for MACD", longTerm, shortTerm)
  }
  twelvePeriodEMA, err := sma.ExponentialMovingAverageSeries(shortTerm, prices)
  if err != nil {
    return nil, err
  }
  twentySixPeriodEMA, err := sma.ExponentialMovingAverageSeries(longTerm, prices)
  if err != nil {
    return nil, err
  }

  macd := make([]float64, len(twentySixPeriodEMA))
  twelvePeriodEMA = twelvePeriodEMA[len(twelvePeriodEMA)-len(macd):]

  for i:= 0; i<len(macd); i++ {
    macd[i] = twelvePeriodEMA[i] - twentySixPeriodEMA[i]
  }
  return macd, nil
}
