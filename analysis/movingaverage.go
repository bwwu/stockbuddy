package main

import (
  "fmt"
  quotepb "stockbuddy/protos/quote_go_proto"
)

func main() {
  x := quotepb.Quote{
    Symbol: "GOOG", 
    Timestamp: 12345, 
    Open: 14.1, 
    High: 14.7,
    Low: 13.9,
    Close: 12.6,
    AdjClose: 7.3, 
    Volume: 12000,
}
  fmt.Printf("%s", x.Symbol)
}
