package moving_average

import (
  "bytes"
  "errors"
  "fmt"
  "stockbuddy/analysis/lib/sma"
  cr "stockbuddy/analysis/moving_average/crossover/crossover_reporter"
  quotepb "stockbuddy/protos/quote_go_proto"
)


// NewMovingAverageCrossoverReporter
func NewMovingAverageCrossoverReporter(shortTerm int, longTerm int, quotes []*quotepb.Quote) (*cr.CrossoverReporter, error) {
  if shortTerm >= longTerm || shortTerm <= 0 {
    return nil, errors.New(fmt.Sprintf("Invalid long(%d) and short(%d) term values for series.", longTerm, shortTerm))
  }
  if len(quotes) < longTerm+1 {
    return nil, errors.New(fmt.Sprintf("Unable to compute N-series SMA with N=%d for series length %d", longTerm, len(quotes)))
  }

  quotes = quotes[len(quotes)-longTerm-1:]
  prices := make([]float64, longTerm+1)
  for i:=0; i<len(prices); i++ {
    prices[i] = quotes[i].Close
  }
  shortMA := sma.SimpleMovingAverageSeries(shortTerm, prices)
  longMA := sma.SimpleMovingAverageSeries(longTerm, prices)

  reporter := cr.NewCrossoverReporter("SMA(12,48)", quotes[0].Symbol, "12-Day SMA", "48-Day SMA", shortMA, longMA)
  return reporter, nil
}

// GetSummaryTable returns an html formatted summary table.
func GetSummaryTable(summaries []*cr.CrossoverReporter) string {
  var b bytes.Buffer
  table :="<table width=\"640\" align=\"center\" border=\"1\">\n"
  //heading := "<tr><th>SYM</th><th>12DMA</th><th>48DMA</th><th>SIGNAL</th></tr>\n"

  b.WriteString(table)
  b.WriteString(summaries[0].GetGenericTableHeader())

  for _, r := range summaries {
   b.WriteString(r.GetTableRow(true))
  }
  b.WriteString("</table>")
  return b.String()
}
