package crossover

import (
  "log"
  "stockbuddy/analysis/constants"
)

// DetectCrossover determines Crossover points of two price series.
// A Bullish crossover occurs when the "cmp" series rises above
// the "ref" series.
func DetectCrossovers(cmp, ref []float64) []constants.Outlook {
  // Canonicalize the series lengths by removing excess from the head
  adjustedLen := min(len(cmp), len(ref))
  cmp = cmp[len(cmp)-adjustedLen:]
  ref = ref[len(ref)-adjustedLen:]

  // Positive vals mean cmp series is leading the ref
  diffs := diffLists(cmp, ref)
  crossovers := make([]constants.Outlook, len(diffs)-1)
  for i:=0; i<len(diffs)-1; i++ {
    // Look for +/- transitions
    prev := diffs[i]
    curr := diffs[i+1]

    if prev*curr <= 0 {
      if curr >= 0 && prev < 0 || curr > 0 && prev == 0 {
        crossovers[i] = constants.Bullish
      } else if curr < 0 && prev >= 0 || curr == 0 && prev > 0{
        crossovers[i] = constants.Bearish
      }
      // Defaults to None. Bug if prev = curr = 0
    }
  }
  return crossovers
}

// Helper functions
func min(x, y int) int {
  if x < y {
    return x
  }
  return y
}

func diffLists(minuend, subtrahend []float64) []float64 {
  if len(minuend) != len(subtrahend) {
    log.Fatalf("Minuend series len (%d) must be equal to that of subtrahend series (%d)", len(minuend), len(subtrahend))
  }
  diff := make([]float64, len(minuend))
  for i:=0; i<len(minuend); i++ {
    diff[i] = minuend[i] - subtrahend[i]
  }
  return diff
}
