load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

package(default_visibility = ["//visibility:public"])

go_library(
  name = "rsi",
  srcs = ["rsi.go"],
  deps = [":sma"],
  importpath = "stockbuddy/analysis/lib/rsi",
)

go_test(
  name = "rsi_test",
  srcs = ["rsi_test.go"],
  deps = [":rsi"],
)

go_library(
  name = "sma",
  srcs = ["sma.go"],
  importpath = "stockbuddy/analysis/lib/sma",
)

go_test(
  name = "sma_test",
  srcs = ["sma_test.go"],
  deps = [":sma"],
)

go_library(
  name = "macd",
  srcs = ["macd.go"],
  importpath = "stockbuddy/analysis/lib/macd",
  deps = [
    "//analysis/lib:sma",
  ],
)

go_test(
  name = "macd_test",
  srcs = ["macd_test.go"],
  deps = [
    ":macd",
    "//analysis/lib:sma",
  ],
)
