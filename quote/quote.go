package quote

type Quote struct {
	Symbol    string
	Timestamp int64
	Open      float64
	High      float64
	Low       float64
	Close     float64
	AdjClose  float64
	Volume    uint64
}
