package constants

// Outlook indicates whether a trend is Bearish/Bullish
type Outlook int

const (
	Bearish Outlook = iota + 1
	Bullish
)

func (r Outlook) String() string {
	switch r {
	case Bearish:
		return "Bearish"
	case Bullish:
		return "Bullish"
	default:
		return "None"
	}
}

type Trend int

const (
	Reversal Trend = iota + 1
	Continuation
	Neither
)

type PriceExtension int

const (
	Overbought PriceExtension = iota + 1
	Oversold
)

func (pe PriceExtension) String() string {
	switch pe {
	case Overbought:
		return "Overbought"
	case Oversold:
		return "Oversold"
	default:
		return "N/A"
	}
}
