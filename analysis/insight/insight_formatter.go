package insight

import (
  "bytes"
  "fmt"
)

func TableFormat(summaries []*AnalyzerSummary) string {
  if len(summaries) == 0 {
    return ""
  }

  heading :=`<table width="640" align="center" border="1">` +
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

