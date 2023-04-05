package percentile

import (
	"log"
	"testing"
)

func TestPercentile(t *testing.T) {
	checkAt(t, 15, 5, testSeries)
	checkAt(t, 20, 30, testSeries)
	checkAt(t, 20, 40, testSeries)
	checkAt(t, 35, 50, testSeries)
	checkAt(t, 50, 100, testSeries)

	checkOf(t, 0, 15, testSeries)
	checkOf(t, 0, 14, testSeries)
	checkOf(t, 20, 16, testSeries)
	checkOf(t, 20, 20, testSeries)
	checkOf(t, 40, 21, testSeries)
	checkOf(t, 40, 35, testSeries)
	checkOf(t, 60, 40, testSeries)
	checkOf(t, 80, 50, testSeries)
	checkOf(t, 100, 51, testSeries)
}

func checkAt(t *testing.T, want, p int, input []int) {
	got, err := At(p, input)
	if err != nil {
		log.Fatal(err)
	}

	if want != got {
		t.Errorf("At(%v, %v) = %v, want %v", p, input, got, want)
	}
}

func checkOf(t *testing.T, want, val int, input []int) {
	got := Of(val, input)

	if want != got {
		t.Errorf("Of(%v, %v) = %v, want %v", val, input, got, want)
	}
}

var testSeries = []int{40, 20, 15, 35, 50}
