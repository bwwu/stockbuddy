package macd_crossover_test

import (
  "log"
  "testing"
  "stockbuddy/analysis/constants"
  macd "stockbuddy/analysis/detectors/macd_crossover"
  pb "stockbuddy/protos/quote_go_proto"
)

func TestBearishCrossover(t *testing.T) {
  d, err := macd.NewMACDDetector(12, 26, 9)
  if err != nil {
    log.Fatal(err)
  }
  crossover, err := d.Process(pricesToQuotes(testSeries))
  if err != nil {
    log.Fatal(err)
  }
  if crossover == nil {
    t.Errorf("MACDDetector.Process(...) = nil")
  }
  got := crossover.Outlook()
  want := constants.Bearish

  if got != want {
    t.Errorf("crossover.Outlook() = %v, want %v", got, want)
  }

  gotSummary := crossover.Summary()
  wantSummary := "MACD(12,26)=-0.47, Signal Line(9)=-0.35, Delta=-0.12"
  if gotSummary != wantSummary {
    t.Errorf("crossover.Summary() = %v, want %v", gotSummary, wantSummary)
  }
}

func pricesToQuotes(prices []float64) []*pb.Quote {
  quotes := make([]*pb.Quote, 0, len(prices))
  for _, price := range prices {
    quotes = append(quotes, &pb.Quote{Close: price})
  }
  return quotes
}

var testSeries = []float64{52.44, 55.93, 56.35, 56.01, 56.52, 55.69, 54.19, 54.37, 53.93, 53.18, 53.57, 52.03, 51.78, 53.23, 54.75, 55.1, 55.93, 56.42, 57.11, 55.86, 56.17, 54.75, 55.4, 56.05, 56.13, 57.41, 57.03, 57.18, 56.08, 56.6, 55.73, 54.73, 54.74, 55.81, 56.48, 56.6, 56.19, 56.34, 57.13, 57.3, 57.95, 58.05, 57.62, 57.21, 57.74, 57.36, 57.73, 57.71, 57.23, 56.62, 56.53, 56.93, 56.47, 55.4, 55.39, 53.25, 51.37, 52.6, 52.34, 53.16, 52.43, 51.54, 52.72, 50.61, 46.25, 46.96, 48.5, 47.93, 48.77, 48.18, 46.61, 47.1, 46.79, 46.87, 47.27, 46.81, 46.5, 47.32, 48.42, 48.84, 48.58, 49.21, 50.03, 49.93, 50.03, 49.96, 49.41, 49.34, 49.19, 49.6, 49.42, 49.12, 49.61, 48.83, 48.84, 49.41, 47.74, 46.56}
