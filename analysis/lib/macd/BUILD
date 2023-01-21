load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

package(default_visibility = ["//visibility:public"])

go_library(
  name = "macd",
  srcs = ["macd.go"],
  importpath = "stockbuddy/analysis/lib/macd",
  deps = [
    "//analysis/lib/sma",
  ],
)

go_test(
  name = "macd_test",
  srcs = ["macd_test.go"],
  deps = [
    ":macd",
    "//analysis/lib/sma",
  ],
)
