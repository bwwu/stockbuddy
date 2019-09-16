// Test moving_average.go
package moving_average_test

import (
  "testing"
  ma "stockbuddy/analysis/moving_average"
  quotepb "stockbuddy/protos/quote_go_proto"
)

func TestNDayMovingAverageWithOffset(t *testing.T) {
  quotes := createQuoteSeriesFromFloats([]float64{1., 2., 3.})
  twoDayMA, _ := ma.NDayMovingAverageWithOffset(2, 0, quotes)
  testEquals(t, "2-Day Moving Average", 2.5, twoDayMA)

  twoDayMAWithOffset, _ := ma.NDayMovingAverageWithOffset(2, 1, quotes)
  testEquals(t, "2-Day Moving Average with offset", 1.5, twoDayMAWithOffset)
}

func createQuoteSeriesFromFloats(prices []float64) []*quotepb.Quote {
  series := make([]*quotepb.Quote, len(prices))
  for i, price := range prices {
    series[i] = &quotepb.Quote{Close: price}
  }
  return series
}

func testEquals(t *testing.T, label string,  expected float64, actual float64) {
  if actual != expected {
    t.Errorf("%s failed. Expected %v, but got %v", label, expected, actual)
  }
}
