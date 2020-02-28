package main

import (
  "bytes"
  "context"
  "log"
  "math"
  "fmt"
  "google.golang.org/grpc"
  "stockbuddy/smtp"
  pb "stockbuddy/protos/quote_go_proto"
)

var indices = []string{
  "^DJI",   // Dow Jones Industrial Average
  "^IXIC",  // Nasdaq
  "^GSPC",  // S&P500
}

var recipients = []string{
  "brandonwu23@gmail.com",
  "anthonywu.ad@gmail.com",
}

func main() {
  conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
  defer conn.Close()
  if err != nil {
    log.Fatal(err.Error())
  }
  client := pb.NewQuoteServiceClient(conn)

  summaries := make([]*volsummary, len(indices))
  for i, index := range indices {
    summaries[i], _ = summarizeVolume(context.Background(), client, index)
  }
  mail(tableFormat(summaries))
}

func mail(body string) {
  subject := "Index Volume Summary"
  intro := "<p>Today's volume vs last 5, last 253 trading days and percent difference</p>\n"
  email := smtp.Email{intro + body, subject, recipients}
  email.Send()
}

type volsummary struct {
  symbol string
  day, avg5day, avg365day uint32
  pdiffv5day, pdiffv365day float64
}

func summarizeVolume(ctx context.Context, client pb.QuoteServiceClient, symbol string) (*volsummary, error) {
  req := &pb.QuoteRequest{Symbol: symbol, Period: 420}
  resp, err := client.ListQuoteHistory(ctx, req)
  if err != nil {
    return nil, err
  }

  volumes := make([]uint32, 0, len(resp.Quotes))
  for _, quote := range resp.Quotes {
    volumes = append(volumes, quote.Volume)
  }

  day := volumes[len(volumes)-1]
  avg5Day := avgvol(5, volumes)
  avg365Day := avgvol(253, volumes)
  pdiffv5Day := pdiff(avg5Day, day)
  pdiffv365Day := pdiff(avg365Day, day)

  return &volsummary{symbol, day, avg5Day, avg365Day, pdiffv5Day, pdiffv365Day}, nil
}

// avgvol calculates avg vol of a list of uint32s, rouding to a uint32.
func avgvol(n int, volumes []uint32) uint32 {
  var sum uint64
  for _, vol := range volumes[len(volumes)-n:] {
    sum += uint64(vol)
  }
  return uint32(math.Round(float64(sum)/float64(n)))
}

// pdiff calculates percentage difference of reference vs a compare
func pdiff(ref, cmp uint32) float64 {
  diff := cmp-ref
  return 100.*float64(diff)/float64(ref)
}


func tableFormat(summaries []*volsummary) string {
  if len(summaries) == 0 {
    return ""
  }

  heading :=`<table width="640" align="center" border="1">` +
    "<tr><th>INDEX</th><th>VOL</th><th>AVG-5DAY</th><th>%DIFFv5DAY</th><th>AVG-253DAY</th><th>%DIFFv253DAY</th></tr>\n"

  var b bytes.Buffer
  b.WriteString(heading)

  for _, s := range summaries {
    var index string
    if s.symbol == "^DJI" {
      index = "DJIA"
    } else if s.symbol == "^IXIC" {
      index = "Nasdaq"
    } else {
      index = "S&P500"
    }

    b.WriteString(
      fmt.Sprintf("<tr><td>%v</td><td>%v</td><td>%v</td><td>%v</td><td>%v</td><td>%v</td></tr>\n",
        index,
        s.day,
        s.avg5day,
        s.pdiffv5day,
        s.avg365day,
        s.pdiffv365day,
      ),
    )
  }

  b.WriteString("</table>")
  return b.String()
}

