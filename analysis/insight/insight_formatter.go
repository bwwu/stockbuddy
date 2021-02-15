package insight

import (
	"bytes"
	"fmt"
)

type row struct {
	symbol     string
	indicators map[string]Indicator
}

func TableFormat(summaries []*AnalyzerSummary) string {
	if len(summaries) == 0 {
		return ""
	}

	heading := `<table width="640" align="center" border="1">` +
		"<tr><th>SYMBOL</th><th>INDICATOR</th><th>SUMMARY</th><th>OUTLOOK</th></tr>\n"

	var b bytes.Buffer
	b.WriteString(heading)

	for _, summary := range summaries {
		for _, ind := range summary.Indicators {
			b.WriteString(
				fmt.Sprintf("<tr><td>%s</td><td>%s</td><td>%s</td><td>%s</td></tr>\n",
					summary.Symbol,
					ind.Name(),
					ind.Summary(),
					ind.Outlook().String(),
				),
			)
		}
	}

	b.WriteString("</table>")
	return b.String()
}

// FormatByIndicator returns a html table with one row per symbol and one column per indicator.
func FormatByIndicator(summaries []*AnalyzerSummary) string {
	if len(summaries) == 0 {
		return ""
	}
	cols := make(map[string]bool)
	var rows []row
	for _, s := range summaries {
		indicators := make(map[string]Indicator)
		for _, i := range s.Indicators {
			// If new column name discovered, add it to the set.
			if _, ok := cols[i.Name()]; !ok {
				cols[i.Name()] = true
			}
			indicators[i.Name()] = i
		}
		rows = append(rows, row{s.Symbol, indicators})
	}
	heading := `<table width="640" align="center" border="1"><tr><th>Symbol</th>`
	var b bytes.Buffer
	b.WriteString(heading)
	for col := range cols {
		b.WriteString(fmt.Sprintf("<th>%s</th>", col))
	}
	b.WriteString("</tr>\n")

	for _, r := range rows {
		b.WriteString(fmt.Sprintf("<tr><td>%s</td>", r.symbol))
		for col := range cols {
			val, ok := r.indicators[col]
			if !ok {
				// If this symbol doesn't have this indicator, print a blank.
				b.WriteString("<td>-</td>")
			} else {
				b.WriteString(fmt.Sprintf("<td>%s</td>", val.Summary()))
			}
		}
		b.WriteString("</tr>\n")
	}
	b.WriteString("</table>")
	return b.String()
}
