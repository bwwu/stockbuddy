package smax

import (
  "fmt"
  "github.com/bwwu/stockbuddy/analysis/insight"
  "github.com/bwwu/stockbuddy/analysis/constants"
  "github.com/bwwu/stockbuddy/analysis/lib/sma"
  "github.com/bwwu/stockbuddy/analysis/lib/crossover"
  "github.com/bwwu/stockbuddy/quote"
)

// SimpleMovingAverageDetector implements IndicatorFactory interface
type SimpleMovingAverageDetector struct {
  shortTerm, longTerm int
}

func NewSimpleMovingAverageDetector(shortTerm, longTerm int) (*SimpleMovingAverageDetector, error) {
  if shortTerm >= longTerm || shortTerm <= 0 {
    return nil, fmt.Errorf("sma_crossover: invalid long(%d) and short(%d) term values for series.", longTerm, shortTerm)
  }

  return &SimpleMovingAverageDetector{
    shortTerm: shortTerm,
    longTerm: longTerm,
  }, nil
}

func (detector *SimpleMovingAverageDetector) Process(quotes []*quote.Quote) (insight.Indicator, error) {
  if len(quotes) < detector.longTerm+1 {
    return nil, fmt.Errorf(
      "sma_crossover: unable to compute N-series SMA with N=%d for series length %d",
      detector.longTerm,
      len(quotes),
    )
  }

  quotes = quotes[len(quotes)-detector.longTerm-1:]
  prices := make([]float64, detector.longTerm+1)
  for i:=0; i<len(prices); i++ {
    prices[i] = quotes[i].Close
  }
  shortMA, err := sma.SimpleMovingAverageSeries(detector.shortTerm, prices)
  if err != nil {
    return nil, err
  }
  longMA, err := sma.SimpleMovingAverageSeries(detector.longTerm, prices)
  if err != nil {
    return nil, err
  }

  crossovers := crossover.DetectCrossovers(shortMA, longMA)
  recentCrossover := crossovers[len(crossovers)-1]
  if recentCrossover == 0 {
    return nil, nil
  }

  return &SimpleMovingAverageCrossover{
    outlook: recentCrossover,
    shortTerm: detector.shortTerm,
    longTerm: detector.longTerm,
    shortMA: shortMA[len(shortMA)-1],
    shortMADelta: shortMA[len(shortMA)-2],
    longMA: longMA[len(longMA)-1],
    longMADelta: longMA[len(longMA)-2],
  }, nil
}

// SimpleMovingAverageCrossover implements Indicator interface
type SimpleMovingAverageCrossover struct {
  outlook constants.Outlook
  shortTerm, longTerm int
  shortMA, shortMADelta, longMA, longMADelta float64
}

func (c SimpleMovingAverageCrossover) Name() string {
  return fmt.Sprintf("MA-%d vs MA-%d", c.shortTerm, c.longTerm)
}

func (c SimpleMovingAverageCrossover) Summary() string {
  summary := "MA(%d)=%s, MA(%d)=%s"
  return fmt.Sprintf(summary, c.shortTerm,
    getDeltaStr(c.shortMA, c.shortMADelta),
    c.longTerm,
    getDeltaStr(c.longMA, c.longMADelta),
  )
}

func (c SimpleMovingAverageCrossover) Outlook() constants.Outlook {
  return c.outlook
}

func (c SimpleMovingAverageCrossover) Trend() constants.Trend {
  return constants.Reversal
}

func getDeltaStr(ref, cmp float64) string {
  delta := ref - cmp
  if delta >= 0 {
    return fmt.Sprintf("%.2f(+%.2f)", ref, delta)
  }
  return fmt.Sprintf("%.2f(%.2f)", ref, delta)
}
