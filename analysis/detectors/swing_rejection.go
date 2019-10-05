/**
 * Swing Rejection is a method of interpretting price reversal through RSI. It
 * involves the following 4 steps for bullish reversal (bearish in paren):
 *  1. RSI falls into oversold (overbought) territory
 *  2. RSI crosses back above 30% (below 70%)
 *  3. RSI forms another dip w/o xing back into oversold (overbought) territory
 *  4. RSI then breaks its most recent high (low)
 */
package detectors

import (
  "fmt"
  "stockbuddy/analysis/constants"
  "stockbuddy/analysis/lib/rsi"
)


/**
 * DetectSwingRejection determines whether there was a swing rejection event
 * in the price series, looking back the num of days specified as "lookback".
 */
func DetectSwingRejection(prices []float64, lookback int) (constants.Outlook, error) {

  rsiSeries, err := rsi.RelativeStrengthIndexSeries(14, prices)
  if err != nil {
    return 0, err
  }
  if len(rsiSeries) < lookback {
    return 0, fmt.Errorf("swing_rejection: price series len %d insufficient for lookback period of %d days", len(prices), lookback)
  }

  lookbackSeries := rsiSeries[len(rsiSeries)-lookback:]

  start, extension := findMostRecentPriceExtension(lookbackSeries)
  if start == -1 {
    return 0, nil
  }

  // Indicates occurence of dip (Step 3).
  var dipConfirmed bool
  extremeRsi := lookbackSeries[start]
  for i:=start+1; i<len(lookbackSeries)-1; i++ {
    if extension == constants.Oversold {
      if lookbackSeries[i] > extremeRsi {
        // Step 2: RSI crosses above 30.
        extremeRsi = lookbackSeries[i]
      } else if lookbackSeries[i] < extremeRsi && lookbackSeries[i] > 30. {
        // Step 3: RSI dips without crossing into oversold.
        dipConfirmed = true
      }
    } else {
      if lookbackSeries[i] < extremeRsi {
        // Step 2: RSI crosses below 70
        extremeRsi = lookbackSeries[i]
      } else if lookbackSeries[i] > extremeRsi && lookbackSeries[i] < 70. {
        // Step 3: RSI rises without crossing into overbought.
        dipConfirmed = true
      }
    }
  }

  if dipConfirmed {
    if extension == constants.Oversold && lookbackSeries[len(lookbackSeries)-1] > extremeRsi {
      // Step 4: RSI surpasses previous high. 
      return constants.Bullish, nil
    }
    if extension == constants.Overbought && lookbackSeries[len(lookbackSeries)-1] < extremeRsi {
      // Step 4: RSI drops below previous low. 
      return constants.Bearish, nil
    }
  }
  return 0, nil
}

// Returns index and type of price extension (over{bought,sold}) of most recent in RSI series.
func findMostRecentPriceExtension(series []float64) (int, constants.PriceExtension) {
  var ext constants.PriceExtension

  for i:=len(series)-1; i>=0; i-- {
    if series[i] < 30. {
      return i, constants.Oversold
    }
    if series[i] > 70. {
      return i, constants.Overbought
    }
  }

  return -1, ext
}
