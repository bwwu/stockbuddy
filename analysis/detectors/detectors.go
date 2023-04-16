package detectors

import (
	"errors"
	"fmt"
	"github.com/bwwu/stockbuddy/analysis/detectors/macdrsi"
	"github.com/bwwu/stockbuddy/analysis/detectors/macdx"
	"github.com/bwwu/stockbuddy/analysis/detectors/smax"
	"github.com/bwwu/stockbuddy/analysis/detectors/swingrejection"
	"github.com/bwwu/stockbuddy/analysis/insight"

)

func GetDefaultDetectors(names []string) ([]insight.Detector, error) {
	var detecs []insight.Detector
	var errs []error

	for _, name := range names {

		if detectorf, ok := defaultDetectors[name]; !ok {
			errs = append(errs, fmt.Errorf("detectors::GetDefaultDetectors: invalid detector '%s'", name))
		} else if detector, err := detectorf(); err != nil {
			errs = append(errs, err)
		} else {
			detecs = append(detecs, detector)
		}
	}

	return detecs, errors.Join(errs...)
}

var defaultDetectors = map[string]func()(insight.Detector, error){
	"macd": macd12_26_9,
	"macd_rsi": macd_rsi,
	"sma": sma12_48,
	"swingrejection": swingrection30_14,
}

func macd12_26_9() (insight.Detector, error) {
	return macdx.NewMACDDetector(12, 26, 9)
}

func sma12_48() (insight.Detector, error) {
	return smax.NewSimpleMovingAverageDetector(12, 48)
}

func swingrection30_14() (insight.Detector, error) {
	return swingrejection.NewSwingRejectionDetector(30, 14), nil
}

func macd_rsi() (insight.Detector, error) {
	return macdrsi.NewMACDxRSIDetector(12, 26, 9, 14, 15)
}