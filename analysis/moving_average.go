package moving_average

import (
  "errors"
  "fmt"
  "log"
  quotepb "stockbuddy/protos/quote_go_proto"
)

type CrossoverType int

const (
    None CrossoverType = iota
    Bearish
    Bullish
)

func (c *CrossoverType)  String() string {
  switch *c {
    case Bearish:
      return "Bearish"
    case Bullish:
      return "Bullish"
    default:
      return "None"
  }
}

type MovingAverageCrossoverSummary struct {
  Symbol string
  ShortTerm int
  LongTerm int
  ShortMA  float64
  ShortMAMinus1 float64
  LongMA float64
  LongMAMinus1 float64
  Crossover CrossoverType
}

// NewMovingAverageCrossoverSummary
func NewMovingAverageCrossoverSummary(shortTerm int, longTerm int, quotes []*quotepb.Quote) (*MovingAverageCrossoverSummary, error) {
  if shortTerm >= longTerm || shortTerm <= 0 {
    return nil, errors.New(fmt.Sprintf("Invalid long(%d) and short(%d) term values for series.", longTerm, shortTerm))
  }
  shortSeries, err := NewMovingAverageSeries(shortTerm, quotes)
  if err != nil {
    return nil, err
  }
  longSeries, err := NewMovingAverageSeries(longTerm, quotes)
  if err != nil {
    return nil, err
  }

  // Can ignore errs for delta <= 1
  shortMA, _ := shortSeries.GetAtDelta(0)
  shortMAMinus1, _ := shortSeries.GetAtDelta(1)
  longMA, _ := longSeries.GetAtDelta(0)
  longMAMinus1, _ := longSeries.GetAtDelta(1)

  // If product of the deltas <= 0, there was a crossover
  delta0 := shortMA - longMA
  delta1 := shortMAMinus1 - longMAMinus1

  crossover := None

  if delta0*delta1 <= 0 {
    log.Printf("%d/%d-MA Crossover detected.\n", shortTerm, longTerm)
    if delta0 >= 0 && delta1 < 0 || delta0 > 0 && delta1 == 0 {
      crossover = Bullish
    } else if delta0 < 0 && delta1 >= 0 || delta0 == 0 && delta1 >0 {
      crossover = Bearish
    }
    // Bug when both are 0. Need to look back further.
  }
  summary := &MovingAverageCrossoverSummary{
    Symbol: quotes[0].Symbol,
    ShortTerm: shortTerm,
    LongTerm: longTerm,
    ShortMA: shortMA,
    ShortMAMinus1: shortMAMinus1,
    LongMA: longMA,
    LongMAMinus1: longMAMinus1,
    Crossover: crossover,
  }
  return summary, nil
}

type MovingAverageSeries struct {
  Term int
  quotes []*quotepb.Quote
}

func (series *MovingAverageSeries) GetAtDelta(delta int) (float64, error) {
  return NDayMovingAverageWithOffset(series.Term, delta, series.quotes)
}

func NewMovingAverageSeries(term int, quotes []*quotepb.Quote) (*MovingAverageSeries, error) {
  if term <= 0 {
    return nil, errors.New("MovingAverageSeries Term must be > 0")
  }
  if len(quotes) < term + 1{
    return nil, errors.New(fmt.Sprintf("MovingAverageSeries Term %d must be less than the quote len %d", term, len(quotes)))
  }
  series := new(MovingAverageSeries)
  series.Term = term
  series.quotes = quotes
  return series, nil
}

// NDayMovingAverageWithOffset calculates N-day moving average for a quote series ordered
// in ascending sequential order (newest quote last). 
// An offset "X" can be used to calculate the N-day moving average X days ago. Should default
// to 0.
func NDayMovingAverageWithOffset(n int, offset int, series []*quotepb.Quote) (float64, error) {
  seriesLen := len(series)
  if seriesLen < n + offset {
    return 0., errors.New(fmt.Sprintf("Series len must be >= %d, but it is %d", n + offset, seriesLen))
  }
  // Take the N last quotes in the series
  nDaySeries := series[seriesLen-n-offset:seriesLen-offset]
  accum := 0.
  for _, quote := range nDaySeries {
    accum += quote.Close
  }
  return accum/float64(n), nil
}
