load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

package(default_visibility = ["//visibility:public"])

go_library(
  name = "macd_crossover",
  srcs = ["macd_crossover.go"],
  importpath = "stockbuddy/analysis/detectors/macdx",
  deps = [
    "//analysis/constants",
    "//analysis/insight",
    "//analysis/lib/macd",
    "//analysis/lib/sma",
    "//protos:quote_go_proto",
    "//analysis/lib/crossover",
  ],
)

go_test(
  name = "macd_crossover_test",
  srcs = ["macd_crossover_test.go"],
  deps = [
    ":macd_crossover",
    "//analysis/constants",
    "//protos:quote_go_proto",
  ],
)
load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")
