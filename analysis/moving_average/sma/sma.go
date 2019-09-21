package sma

import (
  "log"
)

// SimpleMovingAverage calculates the N-period arthmetic average of a price
// series. Assumes prices are ordered in descending order of age (oldest
// first).
func SimpleMovingAverage(n int, prices []float64) float64 {
  if len(prices) < n {
    log.Fatal("Price series length must be >= N")
  }
  sum := 0.
  for _, price := range prices[len(prices)-n:] {
    sum += price
  }
  return sum/float64(n)
}

// ExponentialMovingAverage
func ExponentialMovingAverage(n int, prices []float64) float64 {
  // k = 3.45*(N+1) is recommmended min num of data points for accurate EMA
  // en.wikipedia.org/wiki/Moving_average#Approximating_the_EMA_with_a_limited_number_of_terms
  if float64(len(prices)) < 3.45*float64(n+1) {
    log.Println("WARNING: for an accurate EMA, the price series should have at least 3.45(N+1) data points")
  }
  return exponentialMovingAverageRec(n, prices)
}

func exponentialMovingAverageRec(n int, prices []float64) float64 {
  if len(prices) == 1 {
    return prices[0]
  }
  alpha := 2./float64(n+1)
  return alpha * prices[len(prices)-1] + (1.-alpha)*exponentialMovingAverageRec(n, prices[:len(prices)-1])
}
