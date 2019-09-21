package moving_average

import (
  "bytes"
  "errors"
  "fmt"
  "log"
  "stockbuddy/analysis/moving_average/sma/sma"
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

  if len(quotes) < longTerm+1 {
    return nil, errors.New(fmt.Sprintf("Unable to compute N-series SMA with N=%d for series length %d", longTerm, len(quotes)))
  }
  prices := make([]float64, longTerm+1)
  for i:=0; i<len(prices); i++ {
    prices[i] = quotes[i].Close
  }

  // Can ignore errs for delta <= 1
  shortMA := sma.SimpleMovingAverage(shortTerm, prices)
  shortMAMinus1 := sma.SimpleMovingAverage(shortTerm, prices[:len(prices)-1])
  longMA := sma.SimpleMovingAverage(longTerm, prices)
  longMAMinus1 := sma.SimpleMovingAverage(longTerm, prices[:len(prices)-1])

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

// GetSummaryTable returns an html formatted summary table.
func GetSummaryTable(summaries []*MovingAverageCrossoverSummary) string {
  var b bytes.Buffer
  table :="<table width=\"640\" align=\"center\" border=\"1\">\n"
  heading := "<tr><th>SYM</th><th>12DMA</th><th>12DMAΔ</th><th>48DMA</th><th>48DMAΔ</th><th>SIGNAL</th></tr>\n"

  b.WriteString(table)
  b.WriteString(heading)

  for _, s := range summaries {
   b.WriteString(formatMACrossoverRow(s))
  }
  b.WriteString("</table>")
  return b.String()
}

func formatMACrossoverRow(s *MovingAverageCrossoverSummary) string {
  shortDelta := s.ShortMA - s.ShortMAMinus1
  longDelta := s.LongMA - s.LongMAMinus1

  var signal string
  if s.Crossover == Bullish {
    signal = "BUY"
  } else {
    signal = "SELL"
  }

  template := "<tr><td>%s</td><td>%.2f</td><td>%.2f</td><td>%.2f</td><td>%.2f</td><td>%s</td></tr>\n"
  return fmt.Sprintf(template, s.Symbol, s.ShortMA, shortDelta, s.LongMA, longDelta, signal)
}
