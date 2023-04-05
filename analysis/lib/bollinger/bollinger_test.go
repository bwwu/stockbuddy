package bollinger

import (
	"log"
	"testing"
)

func TestStandardDeviation(t *testing.T) {
	got, err := StandardDeviation(testVals)
	if err != nil {
		log.Fatal(err)
	}
	want := 11.739817553869887
	if want != got {
		t.Errorf("StandardDeviation(%v) = %v, want %v", testVals, got, want)
	}
}

func TestBollingerBands(t *testing.T) {
	bands, err := BollingerBandSeries(20, 2, testPrices)
	if err != nil {
		log.Fatal(err)
	}
	gotLen := len(bands.Upper)

	gotUBand := bands.Upper[gotLen-1]
	wantUBand := 1260.3533842349082
	if gotUBand != wantUBand {
		t.Errorf("BollingerBandSeries(%v).Upper = %v, want %v", testPrices, gotUBand, wantUBand)
	}

	gotLBand := bands.Lower[gotLen-1]
	wantLBand := 1167.9746157650914
	if gotLBand != wantLBand {
		t.Errorf("BollingerBandSeries(%v).Upper = %v, want %v", testPrices, gotLBand, wantLBand)
	}

	gotMA := bands.MA[gotLen-1]
	wantMA := 1214.1639999999998
	if gotMA != wantMA {
		t.Errorf("BollingerBandSeries(%v).Upper = %v, want %v", testPrices, gotMA, wantMA)
	}

	wantLen := 7
	if wantLen != gotLen {
		t.Errorf("len(BollingerBandSeries(%v)) = %v, want %v", testPrices, gotLen, wantLen)
	}
}

var testVals = []float64{587.656616, 594.447937, 585.76178, 579.538879, 586.380066, 573.485474, 575.519897, 574.781921, 575.779175, 566.714111, 568.519104, 573.704895, 575.769226, 562.196472, 570.932495, 559.344299, 542.999207}

var testPrices = []float64{1182.69, 1191.25, 1189.53, 1151.29, 1168.89, 1167.84, 1171.02, 1192.85, 1188.10, 1168.39, 1181.41, 1211.38, 1204.93, 1204.41, 1206.00, 1220.17, 1234.25, 1239.56, 1231.30, 1229.15, 1232.41, 1238.71, 1229.93, 1234.03, 1218.76, 1246.52}
