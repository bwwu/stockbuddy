package constants

// Outlook indicates whether a trend is Bearish/Bullish
type Outlook int

const (
  Bearish Outlook = iota + 1
  Bullish
)

func (r Outlook)  String() string {
  switch r {
    case Bearish:
      return "Bearish"
    case Bullish:
      return "Bullish"
    default:
      return "None"
  }
}
