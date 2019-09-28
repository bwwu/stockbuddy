package crossover_test

import (
  "testing"
  "stockbuddy/analysis/constants"
  "stockbuddy/analysis/moving_average/crossover/crossover"
)

func TestCrossoverSeries(t *testing.T) {
  cmp := []float64{5., 3., 1., 9., 6.}
  ref := []float64{0., 4., 0., 7., 8.}
  actual := crossover.DetectCrossovers(cmp, ref)
  expected := []constants.Outlook{constants.Bearish, constants.Bullish, 0, constants.Bearish}
  testCrossoverSeriesEqual(t, expected, actual)
}

func testCrossoverSeriesEqual(t *testing.T, expected, actual []constants.Outlook) {
  for i:=0; i<len(expected); i++ {
    if expected[i] != actual[i] {
      t.Errorf("Expected %s crossover, but got %s", expected[i].String(), actual[i].String())
    }
  }
}
