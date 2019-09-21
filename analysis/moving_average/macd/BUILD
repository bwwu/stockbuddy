load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

package(default_visibility = ["//visibility:public"])

go_library(
  name = "macd",
  srcs = ["macd.go"],
  importpath = "stockbuddy/analysis/moving_average/macd/macd",
  deps = [
    "//analysis/moving_average/sma:sma",
  ],
)

go_test(
  name = "macd_test",
  srcs = ["macd_test.go"],
  deps = [":macd"],
)
