package detectors_test

import (
  "log"
  "testing"
  "stockbuddy/analysis/constants"
  pb "stockbuddy/protos/quote_go_proto"
  sma "stockbuddy/analysis/detectors/sma_crossover"
)

func TestBearishSMACrossover(t *testing.T) {
  detector,_ := sma.NewSimpleMovingAverageDetector(12, 48)
  quotes := generateQuotes(testBearish)
  hasCrossover, err := detector.Process(quotes)
  if err != nil {
    log.Fatal(err)
  }
  if !hasCrossover {
    t.Error("d.Process(...) = %v, want %v", hasCrossover, true)
  }
  crossover := detector.Get()
  if crossover.Outlook() != constants.Bearish {
    t.Errorf("d.Get().Outlook() = %v, want %v", crossover.Outlook().String(), constants.Bearish.String())
  }

  got := crossover.Summary()
  want := "MA(12)=6.86(-0.19), MA(48)=6.99(-0.04)"
  if got != want {
    t.Errorf("d.Get().Summary() = %v, want %v", got, want)
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
