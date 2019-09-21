package moving_average_test

import (
  "testing"
  ma "stockbuddy/analysis/moving_average/moving_average"
  quotepb "stockbuddy/protos/quote_go_proto"
)

// Crossover tests
func TestBullishCrossover(t *testing.T) {
  quotes := createQuoteSeriesFromFloats([]float64{50.,30.,20.,100.})
  summary, _ := ma.NewMovingAverageCrossoverSummary(2, 3, quotes)
  testCrossoverTypeEquals(t, ma.Bullish, summary.Crossover)
}

func TestBearishCrossover(t *testing.T) {
  quotes := createQuoteSeriesFromFloats([]float64{15.,100.,20.,10.})
  summary, _ := ma.NewMovingAverageCrossoverSummary(2, 3, quotes)
  testCrossoverTypeEquals(t, ma.Bearish, summary.Crossover)
}

func TestNoneCrossover(t *testing.T) {
  quotes := createQuoteSeriesFromFloats([]float64{100.,50.,25.,10.})
  summary, _ := ma.NewMovingAverageCrossoverSummary(2, 3, quotes)
  testCrossoverTypeEquals(t, ma.None, summary.Crossover)
}

// Helper utils
func createQuoteSeriesFromFloats(prices []float64) []*quotepb.Quote {
  series := make([]*quotepb.Quote, len(prices))
  for i, price := range prices {
    series[i] = &quotepb.Quote{Close: price, Symbol: ""}
  }
  return series
}

func testCrossoverTypeEquals(t *testing.T, expected ma.CrossoverType, actual ma.CrossoverType) {
  if actual != expected {
    t.Errorf("Expected Crossover to be %s, but got %s", expected.String(), actual.String())
  }
}
