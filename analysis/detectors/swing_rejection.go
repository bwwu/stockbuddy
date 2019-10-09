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
  "stockbuddy/analysis/insight"
  "stockbuddy/analysis/lib/rsi"
  pb "stockbuddy/protos/quote_go_proto"
)

type SwingRejection struct {
  lookback, period int
  rsi float64
  outlook constants.Outlook
}

func (sr SwingRejection) Name() string {
  return fmt.Sprintf("RSI(%d) %d-Day Swing Rejection", sr.period, sr.lookback)
}

func (sr SwingRejection) Summary() string {
  return fmt.Sprintf("RSI(%d) = %.2f", sr.period, sr.rsi)
}

func (sr SwingRejection) Outlook() constants.Outlook {
  return sr.outlook
}

func (sr SwingRejection) Trend() constants.Trend {
  return constants.Reversal
}

type SwingRejectionDetector struct {
  // Lookback defines how far back in a price series to look for swings.
  // Period defines the N-period smoothing param for calculating RSI.
  lookback, period int
}

func NewSwingRejectionDetector(lookback, period int) (*SwingRejectionDetector) {
  return &SwingRejectionDetector{
    lookback,
    period,
  }
}

func (d *SwingRejectionDetector) Process(quotes []*pb.Quote) (insight.Indicator, error) {
  prices := make([]float64, 0, len(quotes))
  for _, quote := range quotes {
    prices = append(prices, quote.Close)
  }

  rsiSeries, err := rsi.RelativeStrengthIndexSeries(d.period, prices)
  if err != nil {
    return nil, err
  }
  if len(rsiSeries) < d.lookback {
    return nil, fmt.Errorf(
      "swing_rejection: price series len %d insufficient for lookback period of %d days",
      len(prices),
      d.lookback,
    )
  }

  lookbackSeries := rsiSeries[len(rsiSeries)-d.lookback:]

  start, extension := findMostRecentPriceExtension(lookbackSeries)
  if start == -1 {
    return nil, nil
  }

  // Indicates occurence of dip (Step 3).
  var dipConfirmed bool
  extremeRsi := lookbackSeries[start]
  for i:=start+1; i<len(lookbackSeries)-1; i++ {
    if dipConfirmed {
      // Today's RSI must be the first time it crosses the extreme value.
      // If it drops below the previous low, or above the high before today,
      // then do not detect a swing rejection for today.
      if extension == constants.Oversold && lookbackSeries[i] > extremeRsi {
        return nil, nil
      }
      if extension == constants.Overbought && lookbackSeries[i] < extremeRsi {
        return nil, nil
      }
    } else if extension == constants.Oversold {
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
      return &SwingRejection{
        d.lookback,
        d.period,
        lookbackSeries[len(lookbackSeries)-1],
        constants.Bullish,
      }, nil
    }
    if extension == constants.Overbought && lookbackSeries[len(lookbackSeries)-1] < extremeRsi {
      // Step 4: RSI drops below previous low. 
      return &SwingRejection{
        d.lookback,
        d.period,
        lookbackSeries[len(lookbackSeries)-1],
        constants.Bearish,
       }, nil
    }
  }
  return nil, nil
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
