package crossover_test

import (
  "testing"
  "stockbuddy/analysis/moving_average/crossover/crossover"
)

func TestCrossoverSeries(t *testing.T) {
  cmp := []float64{5., 3., 1., 9., 6.}
  ref := []float64{0., 4., 0., 7., 8.}
  actual := crossover.DetectCrossovers(cmp, ref)
  expected := []crossover.CrossoverType{crossover.Bearish, crossover.Bullish, crossover.None, crossover.Bearish}
  testCrossoverSeriesEqual(t, expected, actual)
}

func testCrossoverSeriesEqual(t *testing.T, expected, actual []crossover.CrossoverType) {
  for i:=0; i<len(expected); i++ {
    if expected[i] != actual[i] {
      t.Errorf("Expected %s crossover, but got %s", expected[i].String(), actual[i].String())
    }
  }
}
