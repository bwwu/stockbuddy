package percentile

import (
  "fmt"
  "sort"
  "math"
)

// At returns the value in a list which is at a certain percentile.
// For example, `At(90, values)` returns the value at 90th percentile.
func At(percentile int, values []int) (int, error) {
	if percentile < 0 || percentile > 100 {
		return 0, fmt.Errorf("percentile: invalid percentile %d", percentile)
	}
	sort.Ints(values)
	ordinal := int(math.Ceil(float64(percentile)/100.0*float64(len(values)))) - 1
	if ordinal < 0 {
		ordinal = 0
	}
	return values[ordinal], nil
}

// Of returns the percentile at a certain value, between 0 and 100.
// For example, `Of(997, values)` returns the percentile of 997 in a list of vals.
func Of(val int, values []int) int {
	sort.Ints(values)
	firstGTE := sort.Search(len(values), func (i int) bool { return values[i] >= val })
	return firstGTE*100/len(values)
}
