package moving_average_test

import (
  "log"
  "testing"
  "stockbuddy/analysis/lib/sma"
)

// Simple Moving Average (SMA) tests
func TestSimpleMovingAverage50(t *testing.T) {
  period := 50
  want := 1187.2912033799998
  got, err := sma.SimpleMovingAverage(period, testSeries)
  if err != nil {
    log.Fatal(err)
  }
  if want != got {
    t.Errorf("sma.SimpleMovingAverage(%v,%v) = %v, want %v", period, testSeries, got, want)
  }
}

func TestSimpleMovingAverage200(t *testing.T) {
  period := 200
  want := 1140.4932497750003
  got, err := sma.SimpleMovingAverage(period, testSeries)
  if err != nil {
    log.Fatal(err)
  }

  if want != got {
    t.Errorf("sma.SimpleMovingAverage(%v,%v) = %v, want %v", period, testSeries, got, want)
  }
}

func TestSimpleMovingAverageSeries50(t *testing.T) {
  subSeries := testSeries[len(testSeries)-52:]
  period := 50
  smaSeries, err := sma.SimpleMovingAverageSeries(period, subSeries)
  if err != nil {
    log.Fatal(err)
  }

  testSliceLen(t, 3, smaSeries)
  if len(smaSeries) != 3 {
    t.Errorf("len(sma.SimpleMovingAverageSeries(%v,%v)) = %v, want %v", period, testSeries, len(smaSeries), 3)
  }
  wantSMA2 := 1187.2912033799998
  gotSMA2 := smaSeries[2]
  if wantSMA2 != gotSMA2 {
    t.Errorf("sma.SimpleMovingAverageSeries(%v,%v) = %v, want %v", period, testSeries, gotSMA2, wantSMA2)
  }

  wantSMA, err := sma.SimpleMovingAverage(period, testSeries[:len(testSeries)-1])
  if err != nil {
    log.Fatal(err)
  }
  gotSMA := smaSeries[1]
  if wantSMA != gotSMA {
    t.Errorf("sma.SimpleMovingAverageSeries(%v,%v) = %v, want %v", period, testSeries, gotSMA, wantSMA)
  }
}

// Exponential Moving Average (EMA) tests
func TestExponentialMovingAverage10(t *testing.T) {
  period := 10
  want := 1224.7593918576256
  got := sma.ExponentialMovingAverage(period, testSeries)

  if want != got {
    t.Errorf("sma.ExponentialMovingAverage(%v,%v) = %v, want %v", period, testSeries, got, want)
  }
}

func TestExponentialMovingAverage30(t *testing.T) {
  period := 30
  want := 1202.3654628306713
  got := sma.ExponentialMovingAverage(period, testSeries)

  if want != got {
    t.Errorf("sma.ExponentialMovingAverage(%v,%v) = %v, want %v", period, testSeries, got, want)
  }
}

// Helper functions
func testFloatEquals(t *testing.T,  expected float64, actual float64) {
  if actual != expected {
    t.Errorf("Expected %v, but got %v", expected, actual)
  }
}

func testSliceLen(t *testing.T, expected int, got []float64) {
  actual := len(got)
  if actual != expected {
    t.Errorf("Expected slice len %d, but got %d", expected, actual)
  }
}

var testSeries = []float64{1106.430054, 1050.819946, 1068.72998, 1036.579956, 1039.550049, 1051.75, 1063.680054, 1061.900024, 1042.099976, 1016.530029, 1028.709961, 1023.01001, 1009.409973, 979.539978, 976.219971, 1039.459961, 1043.880005, 1037.079956, 1035.609985, 1045.849976, 1016.059998, 1070.709961, 1068.390015, 1076.280029, 1074.660034, 1070.329956, 1057.189941, 1044.689941, 1077.150024, 1080.969971, 1089.900024, 1098.26001, 1070.52002, 1075.569946, 1073.900024, 1090.98999, 1070.079956, 1060.619995, 1089.060059, 1116.369995, 1110.75, 1132.800049, 1145.98999, 1115.22998, 1098.709961, 1095.060059, 1095.01001, 1121.369995, 1120.160034, 1121.670044, 1113.650024, 1118.560059, 1113.800049, 1096.969971, 1110.369995, 1109.400024, 1115.130005, 1116.050049, 1119.920044, 1140.98999, 1147.800049, 1162.030029, 1157.859985, 1143.300049, 1142.319946, 1175.76001, 1193.199951, 1193.319946, 1185.550049, 1184.459961, 1184.26001, 1198.849976, 1223.969971, 1231.540039, 1205.5, 1193, 1184.619995, 1173.02002, 1168.48999, 1173.310059, 1194.430054, 1200.48999, 1205.920044, 1215, 1207.150024, 1203.839966, 1197.25, 1202.160034, 1204.619995, 1217.869995, 1221.099976, 1227.130005, 1236.339966, 1236.369995, 1248.839966, 1264.550049, 1256, 1263.449951, 1272.180054, 1287.579956, 1188.47998, 1168.079956, 1162.609985, 1185.400024, 1189.390015, 1174.099976, 1166.27002, 1162.380005, 1164.27002, 1132.030029, 1120.439941, 1164.209961, 1178.97998, 1162.300049, 1138.849976, 1149.630005, 1151.420044, 1140.77002, 1133.469971, 1134.150024, 1116.459961, 1117.949951, 1103.630005, 1036.22998, 1053.050049, 1042.219971, 1044.339966, 1066.040039, 1080.380005, 1078.719971, 1077.030029, 1088.77002, 1085.349976, 1092.5, 1103.599976, 1102.329956, 1111.420044, 1121.880005, 1115.52002, 1086.349976, 1079.800049, 1076.01001, 1080.910034, 1097.949951, 1111.25, 1121.579956, 1131.589966, 1116.349976, 1124.829956, 1140.47998, 1144.209961, 1144.900024, 1150.339966, 1153.579956, 1146.349976, 1146.329956, 1130.099976, 1138.069946, 1146.209961, 1137.810059, 1132.119995, 1250.410034, 1239.410034, 1225.140015, 1216.680054, 1209.01001, 1193.98999, 1152.319946, 1169.949951, 1173.98999, 1204.800049, 1188.01001, 1174.709961, 1197.27002, 1164.290039, 1167.26001, 1177.599976, 1198.449951, 1182.689941, 1191.25, 1189.530029, 1151.290039, 1168.890015, 1167.839966, 1171.02002, 1192.849976, 1188.099976, 1168.390015, 1181.410034, 1211.380005, 1204.930054, 1204.410034, 1206, 1220.170044, 1234.25, 1239.560059, 1231.300049, 1229.150024, 1232.410034, 1238.71, 1229.93}

