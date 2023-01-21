load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

package(default_visibility = ["//visibility:public"])

go_library(
  name = "sma_crossover",
  srcs = ["sma_crossover.go"],
  importpath = "stockbuddy/analysis/detectors/smax",
  deps = [
    "//analysis/constants",
    "//analysis/insight",
    "//analysis/lib/sma",
    "//protos:quote_go_proto",
    "//analysis/lib/crossover",
  ],
)

go_test(
  name = "sma_crossover_test",
  srcs = ["sma_crossover_test.go"],
  deps = [
    ":sma_crossover",
    "//analysis/constants",
    "//protos:quote_go_proto",
  ],
)
