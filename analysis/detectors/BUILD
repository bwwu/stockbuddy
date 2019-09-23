load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

package(default_visibility = ["//visibility:public"])

go_library(
  name = "macd_crossover",
  srcs = ["macd_crossover.go"],
  importpath = "stockbuddy/analysis/detectors/macd_crossover",
  deps = [
    "//analysis/lib:macd",
    "//analysis/lib:sma",
    "//analysis/moving_average/crossover:crossover_reporter",
    "//protos:quote_go_proto",
  ],
)
