package bollinger

import (
	"github.com/bwwu/stockbuddy/analysis/lib/sma"
	"math"
)

type BollingerBands struct {
	Upper, Lower, MA []float64
}

func StandardDeviation(values []float64) (float64, error) {
	n := len(values)
	mean, err := sma.SimpleMovingAverage(n, values)
	if err != nil {
		// Should never happen since N = len(values).
		return 0., err
	}

	var sum float64
	for _, val := range values {
		sum += math.Pow(val-mean, 2)
	}
	return math.Sqrt(sum / float64(n)), nil
}

func StandardDeviationSeries(n int, values []float64) ([]float64, error) {
	stdevs := make([]float64, len(values)-n+1)
	for i := 0; i < len(stdevs); i++ {
		stdev, err := StandardDeviation(values[i : i+n])
		if err != nil {
			return nil, err
		}
		stdevs[i] = stdev
	}
	return stdevs, nil
}

// BollingerBandSeries returns daily Bollinger band values given a price series.
// Requires >= N+1 data points in the series.
func BollingerBandSeries(n, k int, prices []float64) (*BollingerBands, error) {
	// N-series Moving averages
	nmas, err := sma.SimpleMovingAverageSeries(n, prices)
	if err != nil {
		return nil, err
	}
	stdevs, err := StandardDeviationSeries(n, prices)
	if err != nil {
		return nil, err
	}
	// UpperBand = MA(N) + K*σ(N).
	lband := make([]float64, len(stdevs))
	uband := make([]float64, len(stdevs))

	for i := 0; i < len(nmas); i++ {
		lband[i] = nmas[i] - float64(k)*stdevs[i]
		uband[i] = nmas[i] + float64(k)*stdevs[i]
	}
	return &BollingerBands{
		Upper: uband,
		Lower: lband,
		MA:    nmas,
	}, nil
}
