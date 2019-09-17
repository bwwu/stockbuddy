package moving_average_test

import (
  "testing"
  ma "stockbuddy/analysis/moving_average"
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

// Moving avg tests
func TestNDayMovingAverageWithOffset(t *testing.T) {
  quotes := createQuoteSeriesFromFloats([]float64{1., 2., 3.})
  twoDayMA, _ := ma.NDayMovingAverageWithOffset(2, 0, quotes)
  testFloatEquals(t, "2-Day Moving Average", 2.5, twoDayMA)

  twoDayMAWithOffset, _ := ma.NDayMovingAverageWithOffset(2, 1, quotes)
  testFloatEquals(t, "2-Day Moving Average with offset", 1.5, twoDayMAWithOffset)
}

// Helper utilities
func createQuoteSeriesFromFloats(prices []float64) []*quotepb.Quote {
  series := make([]*quotepb.Quote, len(prices))
  for i, price := range prices {
    series[i] = &quotepb.Quote{Close: price, Symbol: ""}
  }
  return series
}

func testFloatEquals(t *testing.T, label string,  expected float64, actual float64) {
  if actual != expected {
    t.Errorf("%s failed. Expected %v, but got %v", label, expected, actual)
  }
}

func testCrossoverTypeEquals(t *testing.T, expected ma.CrossoverType, actual ma.CrossoverType) {
  if actual != expected {
    t.Errorf("Expected Crossover to be %s, but got %s", expected.String(), actual.String())
  }
}
