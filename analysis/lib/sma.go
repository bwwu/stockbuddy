package sma

import (
  "fmt"
  "log"
)

/**
 * SimpleMovingAverage calculates the N-period arthmetic mean of a price
 * series. Assumes prices are ordered in descending order of age (oldest
 * first).
 */
func SimpleMovingAverage(n int, prices []float64) (float64, error) {
  if n < 1 {
   return 0., fmt.Errorf("sma: invalid value for N=%v", n)
  }
  if len(prices) < n {
    return 0., fmt.Errorf("sma: price series length (%d) must be >= N (%d)", len(prices), n)
  }
  sum := 0.
  for _, price := range prices[len(prices)-n:] {
    sum += price
  }
  return sum/float64(n), nil
}

func SimpleMovingAverageSeries(n int, prices []float64) ([]float64, error) {
  series := make([]float64, len(prices)-n+1)
  for i:=0; i<len(series); i++ {
    ma, err := SimpleMovingAverage(n, prices[:len(prices)-i])
    if err != nil {
      return nil, err
    }
    series[len(series)-i-1] = ma
  }
  return series, nil
}

/**
 * ExponentialMovingAverage calculates the weighted average of a price series.
 * Using the formula: EMA(t) = A*Price(t) + (A-1)*EMA(t-1)
 * Where A = 2/(N+1). EMA(0) = Price(0).
 */
func ExponentialMovingAverage(n int, prices []float64) float64 {
  series := ExponentialMovingAverageSeries(n, prices)
  return series[len(series)-1]
}

// ExponentialMovingAverageSeries returns a daily series of EMA given a price
// series.
func ExponentialMovingAverageSeries(n int, prices []float64) []float64 {
  // k = 3.45*(N+1) is recommmended min num of data points for accurate EMA
  // en.wikipedia.org/wiki/Moving_average#Approximating_the_EMA_with_a_limited_number_of_terms
  if float64(len(prices)) < 3.45*float64(n+1) {
    log.Println("WARNING: for an accurate EMA, the price series should have at least 3.45(N+1) data points")
  }
  series := make([]float64, len(prices))

  series[0] = prices[0]
  alpha := 2./float64(n+1)

  for i:=1; i<len(series); i++ {
    series[i] = alpha*prices[i] + (1.-alpha)*series[i-1]
  }
  return series
}
