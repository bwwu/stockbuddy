package constants

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

type PriceExtension int

const (
	Overbought PriceExtension = iota + 1
	Oversold
)
