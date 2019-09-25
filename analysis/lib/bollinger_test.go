package bollinger_test

import (
	"testing"
	"stockbuddy/analysis/lib/bollinger"
)

func TestStandardDeviation(t *testing.T) {
	stdev := bollinger.StandardDeviation(testVals)
	testFloatEquals(t, 11.739817553869887, stdev)
}

func TestBollingerBands(t *testing.T) {
	bands := bollinger.BollingerBandSeries(20, 2, testPrices)
	seriesLen := len(bands.Upper)

	testFloatEquals(t, 1260.3533842349082, bands.Upper[seriesLen-1])
	testFloatEquals(t, 1167.9746157650914, bands.Lower[seriesLen-1])
	testFloatEquals(t, 1214.1639999999998, bands.MA[seriesLen-1])
	testIntEquals(t, 7, seriesLen)
}

func testFloatEquals(t *testing.T,  expected float64, actual float64) {
  if actual != expected {
    t.Errorf("Expected %v, but got %v", expected, actual)
  }
}

func testIntEquals(t *testing.T,  expected int, actual int) {
  if actual != expected {
    t.Errorf("Expected %v, but got %v", expected, actual)
  }
}

var testVals = []float64{587.656616, 594.447937, 585.76178, 579.538879, 586.380066, 573.485474, 575.519897, 574.781921, 575.779175, 566.714111, 568.519104, 573.704895, 575.769226, 562.196472, 570.932495, 559.344299, 542.999207}

var testPrices = []float64{1182.69, 1191.25, 1189.53, 1151.29, 1168.89, 1167.84, 1171.02, 1192.85, 1188.10, 1168.39, 1181.41, 1211.38, 1204.93, 1204.41, 1206.00, 1220.17, 1234.25, 1239.56, 1231.30, 1229.15, 1232.41, 1238.71, 1229.93, 1234.03, 1218.76, 1246.52}
