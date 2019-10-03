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


go_library(
  name = "macd_crossover2",
  srcs = ["macd_crossover2.go"],
  importpath = "stockbuddy/analysis/detectors/macd_crossover2",
  deps = [
    "//analysis:constants",
    "//analysis:insight",
    "//analysis/lib:macd",
    "//analysis/lib:sma",
    "//analysis/moving_average/crossover",
    "//protos:quote_go_proto",
  ],
)

go_test(
  name = "macd_crossover_test",
  srcs = ["macd_crossover_test.go"],
  deps = [
    ":macd_crossover2",
    "//analysis:constants",
    "//protos:quote_go_proto",
  ],
)

go_library(
  name = "sma_crossover",
  srcs = ["sma_crossover.go"],
  importpath = "stockbuddy/analysis/detectors/sma_crossover",
  deps = [
    "//analysis:constants",
    "//analysis:insight",
    "//analysis/lib:sma",
    "//analysis/moving_average/crossover",
    "//protos:quote_go_proto",
  ],
)

go_test(
  name = "sma_crossover_test",
  srcs = ["sma_crossover_test.go"],
  deps = [
    ":sma_crossover",
    "//analysis:constants",
    "//protos:quote_go_proto",
  ],
)
