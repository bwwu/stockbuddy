package detectors

import (
  "stockbuddy/analysis/lib/macd"
  "stockbuddy/analysis/lib/sma"
  cr "stockbuddy/analysis/moving_average/crossover/crossover_reporter"
  quotepb "stockbuddy/protos/quote_go_proto"
)


func DetectMovingAverageConvergenceDivergenceCrossover(quotes []*quotepb.Quote) (*cr.CrossoverReporter, error) {
  prices := make([]float64, len(quotes))
  for i, quote := range quotes {
    prices[i] = quote.Close
  }

  macdSeries := macd.MovingAverageConvergenceDivergenceSeries(prices)
  signalLineSeries := sma.ExponentialMovingAverageSeries(9, macdSeries)

  return cr.NewCrossoverReporter("MACD(12,26,9)", quotes[0].Symbol, "MACD(12,26)", "Signal Line", macdSeries, signalLineSeries), nil
}
