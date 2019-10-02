package bollinger

import (
	"math"
	"stockbuddy/analysis/lib/sma"
)

type BollingerBands struct {
	Upper, Lower, MA []float64
}

func StandardDeviation(values []float64) float64 {
	n := len(values)
	mean := sma.SimpleMovingAverage(n, values)

	var sum float64
	for _,val := range values {
		sum += math.Pow(val - mean, 2)
	}
	return math.Sqrt(sum/float64(n))
}

func StandardDeviationSeries(n int, values []float64) []float64 {
	stdevs := make([]float64, len(values)-n+1)
	for i:=0; i<len(stdevs); i++ {
		stdevs[i] = StandardDeviation(values[i:i+n])
	}
	return stdevs
}

// BollingerBandSeries returns daily Bollinger band values given a price series.
// Requires >= N+1 data points in the series.
func BollingerBandSeries(n, k int, prices []float64) BollingerBands {
	// N-series Moving averages
	nmas := sma.SimpleMovingAverageSeries(n, prices)
	stdevs := StandardDeviationSeries(n, prices)
	// UpperBand = MA(N) + K*σ(N).
	lband := make([]float64, len(stdevs))
	uband := make([]float64, len(stdevs))

	for i:=0; i<len(nmas); i++ {
		lband[i] = nmas[i] - float64(k)*stdevs[i]
		uband[i] = nmas[i] + float64(k)*stdevs[i]
	}
	return BollingerBands{
		Upper: uband,
		Lower: lband,
		MA: nmas,
	}
}
