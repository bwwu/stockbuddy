package detectors_test

import (
  "testing"

  "stockbuddy/analysis/constants"
  pb "stockbuddy/protos/quote_go_proto"
  sma "stockbuddy/analysis/detectors/sma_crossover"
)

func TestBearishSMACrossover(t *testing.T) {
  detector,_ := sma.NewSimpleMovingAverageDetector(12, 48)
  quotes := generateQuotes(testBearish)
  hasCrossover,_ := detector.Process(quotes)

  if !hasCrossover {
    t.Error("Expected a crossover")
  }
  crossover := detector.Get()
  if crossover.Outlook() != constants.Bearish {
    t.Errorf("Expected Bearish outlook, but got %s", crossover.Outlook().String())
  }
}

func generateQuotes(prices []float64) []*pb.Quote {
  quotes := make([]*pb.Quote, len(prices))
  for i,price := range prices {
    quotes[i] = &pb.Quote{Close: price}
  }
  return quotes
}

var testBearish []float64 = []float64{7.08, 7.12, 7.23, 7.3, 7.22, 7.37, 7.34, 7.43, 7.3, 7.27, 7.14, 6.85, 6.78, 6.83, 7.06, 6.96, 6.84, 7., 6.75, 6.68, 6.85, 6.96, 6.89, 7.09, 7.04, 6.98, 6.86, 6.73, 6.81, 6.96, 6.87, 6.76, 6.91, 7.28, 7.19, 6.98, 7.15, 7.51, 7.41, 7.47, 7.59, 7.6, 7.52, 7.57, 7.54, 7.51, 5.81, 5.57, 5.45, 5.26}
