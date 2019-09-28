package crossover_reporter

import (
  "bytes"
  "fmt"
  "log"
  "stockbuddy/analysis/constants"
  "stockbuddy/analysis/moving_average/crossover/crossover"
)

type CrossoverReporter struct {
  Name, Symbol, SeriesAName, SeriesBName string
  SeriesA, SeriesB []float64
  Crossovers []constants.Outlook
}

func NewCrossoverReporter(cname, symbol, snameA, snameB string, sA, sB []float64) *CrossoverReporter {
  reporter := &CrossoverReporter{
    Name: cname,
    Symbol: symbol,
    SeriesAName: snameA,
    SeriesBName: snameB,
    SeriesA: sA,
    SeriesB: sB,
    Crossovers: crossover.DetectCrossovers(sA, sB),
  }
  return reporter
}

// GetCrossover returns most recent series value for Outlook
func (reporter *CrossoverReporter) GetCrossover() constants.Outlook {
  return reporter.Crossovers[len(reporter.Crossovers)-1]
}

// GetTableRowWithLabel returns the html for the row. Set withLabel = True to
// include the name of the Crossover in the first column, e.g. "MACD(12,26,9)"
func (reporter *CrossoverReporter) GetTableRow(withLabel bool) string {
  rowValues := []string{
    reporter.Name,
    reporter.Symbol,
    formatValueWithDelta(reporter.SeriesA),
    formatValueWithDelta(reporter.SeriesB),
    reporter.GetCrossover().String(),
  }
  if withLabel {
    return getRow(false, rowValues)
  }
  return getRow(false, rowValues[1:])
}

func (reporter *CrossoverReporter) GetTableHeader() string {
  headerValues := []string{
    "Symbol",
    reporter.SeriesAName,
    reporter.SeriesBName,
    "Signal",
  }
  return getRow(true, headerValues)
}

func (reporter *CrossoverReporter) GetGenericTableHeader() string {
  headerValues := []string{
    "Indicator",
    "Symbol",
    "Series A",
    "Series B",
    "Signal",
  }
  return getRow(true, headerValues)
}

func getRow(isHeader bool, values []string) string {
  tag := "td"
  if isHeader {
    tag = "td"
  }
  var b bytes.Buffer
  b.WriteString("<tr>")
  for _, val := range values {
   b.WriteString(fmt.Sprintf("<%s>%s</%s>", tag, val, tag))
  }
  b.WriteString("</tr>")
  return b.String()
}

func formatValueWithDelta(values []float64) string {
  if len(values) < 2 {
    log.Fatal("Calculating delta requires length > 1")
  }
  curr := values[len(values)-1]
  prev := values[len(values)-2]
  delta := curr - prev
  if (delta >= 0) {
    // Prepend a plus symbol
    return fmt.Sprintf("%.2f (+%.2f)", curr, delta)
  }
  return fmt.Sprintf("%.2f (%.2f)", curr, delta)
}
