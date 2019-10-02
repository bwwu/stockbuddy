package rsi_test

import (
  "log"
  "testing"
  "stockbuddy/analysis/lib/rsi"
)

func TestRsi14(t *testing.T) {
  want := 61.187409242787304
  got, err := rsi.RelativeStrengthIndex(14, testPrices)
  if err != nil {
    log.Fatal(err)
  }
  if got != want {
    t.Errorf("rsi.RelativeStrengthIndex(%v, %v) = %v, want %v", 14, testPrices, got, want)
  }
}

var testPrices = []float64{1140.47998, 1144.209961, 1144.900024, 1150.339966, 1153.579956, 1146.349976, 1146.329956, 1130.099976, 1138.069946, 1146.209961, 1137.810059, 1132.119995, 1250.410034, 1239.410034, 1225.140015, 1216.680054, 1209.01001, 1193.98999, 1152.319946, 1169.949951, 1173.98999, 1204.800049, 1188.01001, 1174.709961, 1197.27002, 1164.290039, 1167.26001, 1177.599976, 1198.449951, 1182.689941, 1191.25, 1189.530029, 1151.290039, 1168.890015, 1167.839966, 1171.02002, 1192.849976, 1188.099976, 1168.390015, 1181.410034, 1211.380005, 1204.930054, 1204.410034, 1206, 1220.170044, 1234.25, 1239.560059, 1231.300049, 1229.150024, 1232.410034}

