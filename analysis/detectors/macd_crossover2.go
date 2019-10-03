package detectors

import (
  "fmt"
  "stockbuddy/analysis/constants"
  "stockbuddy/analysis/insight"
  "stockbuddy/analysis/lib/macd"
  "stockbuddy/analysis/lib/sma"
  "stockbuddy/analysis/moving_average/crossover/crossover"
  pb "stockbuddy/protos/quote_go_proto"
)

// MACDDetector
type MACDDetector struct {
  shortTerm, longTerm, signalTerm int
  crossover *MACDCrossover
}

func NewMACDDetector(shortTerm, longTerm, signalTerm int) (*MACDDetector, error) {
  if shortTerm >= longTerm || shortTerm <= 0 {
    return nil, fmt.Errorf(
      "macd_crossover: invalid long(%d) and short(%d) term values for MACD",
      longTerm,
      shortTerm,
    )
  }
  return &MACDDetector{
    shortTerm: shortTerm,
    longTerm: longTerm,
    signalTerm: signalTerm,
  }, nil
}

func (d *MACDDetector) Process(quotes []*pb.Quote) (bool, error) {
  // Collect closing prices.
  prices := make([]float64, 0, len(quotes))
  for _, quote := range quotes {
    prices = append(prices, quote.Close)
  }

  macdSeries, err := macd.MovingAverageConvergenceDivergenceSeries(d.shortTerm, d.longTerm, prices)
  if err != nil {
    return false, err
  }
  signalLine, err := sma.ExponentialMovingAverageSeries(d.signalTerm, macdSeries)
  if err != nil {
    return false, err
  }

  crossovers := crossover.DetectCrossovers(macdSeries, signalLine)
  recentCrossover := crossovers[len(crossovers)-1]
  if recentCrossover == 0 {
    return false, nil
  }
  d.crossover = &MACDCrossover {
    shortTerm: d.shortTerm,
    longTerm: d.longTerm,
    signalTerm: d.signalTerm,
    macd: macdSeries[len(macdSeries)-1],
    signalLine: signalLine[len(signalLine)-1],
    outlook: recentCrossover,
  }
  return true, nil
}

func (d *MACDDetector) Get() insight.Indicator {
  return d.crossover
}

// MACDCrossover
type MACDCrossover struct {
  shortTerm, longTerm, signalTerm int
  macd, signalLine float64
  outlook constants.Outlook
}

func (c MACDCrossover) Name() string {
  return fmt.Sprintf("MACD(%d,%d,%d)", c.shortTerm, c.longTerm, c.signalTerm)
}

func (c MACDCrossover) Summary() string {
  return fmt.Sprintf(
    "MACD(%d,%d)=%.2f, Signal Line(%d)=%.2f, Delta=%.2f",
    c.shortTerm,
    c.longTerm,
    c.macd,
    c.signalTerm,
    c.signalLine,
    c.macd-c.signalLine,
  )
}

func (c MACDCrossover) Outlook() constants.Outlook {
  return c.outlook
}

func (c MACDCrossover) Trend() constants.Trend {
  return constants.Reversal
}
