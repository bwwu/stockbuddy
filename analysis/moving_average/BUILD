load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

package(default_visibility = ["//visibility:public"])

go_library(
  name = "moving_average",
  srcs = ["moving_average.go"],
  importpath = "stockbuddy/analysis/moving_average/moving_average",
  deps = [
    "//analysis/lib:sma",
    "//analysis/moving_average/crossover:crossover_reporter",
    "//protos:quote_go_proto",
  ],
)

go_test(
  name = "moving_average_test",
  srcs = ["moving_average_test.go"],
  deps = [
    ":moving_average",
    "//protos:quote_go_proto",
  ],
)
