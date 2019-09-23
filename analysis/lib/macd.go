package macd

import (
  "stockbuddy/analysis/lib/sma"
)

// MovingAverageConvergenceDivergenceSeries calculates the difference between
// 12-Period and 26-Period EMA. Generally compared against a Signal line 
// generated from a 9-Day EMA of the MACD
func MovingAverageConvergenceDivergenceSeries(prices []float64) []float64 {
  twelvePeriodEMA := sma.ExponentialMovingAverageSeries(12, prices)
  twentySixPeriodEMA := sma.ExponentialMovingAverageSeries(26, prices)

  macd := make([]float64, len(twentySixPeriodEMA))
  twelvePeriodEMA = twelvePeriodEMA[len(twelvePeriodEMA)-len(macd):]

  for i:= 0; i<len(macd); i++ {
    macd[i] = twelvePeriodEMA[i] - twentySixPeriodEMA[i]
  }
  return macd
}
