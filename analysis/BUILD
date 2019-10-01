load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
  name = "stocks_to_watch",
  importpath = "stockbuddy/analysis/stocks_to_watch",
  srcs = ["stocks_to_watch.go"],
)

go_library(
  name = "constants",
  srcs = ["constants.go"],
  importpath = "stockbuddy/analysis/constants",
  visibility = ["//visibility:public"],
)

go_library(
  name = "insight",
  srcs = [
    "insight.go",
  ],
  importpath = "stockbuddy/analysis/insight",
  visibility = ["//visibility:public"],
  deps = [
    "//analysis:constants",
    "//protos:quote_go_proto",
  ],
)

go_binary(
  name = "main",
  srcs = ["main.go"],
  deps = [
    ":stocks_to_watch",
    "//analysis/moving_average:moving_average",
    "//analysis/moving_average/crossover:crossover_reporter",
    "//analysis/detectors:macd_crossover",
    "//protos:quote_go_proto",
    "//smtp:sendmail",
    "@org_golang_google_grpc//:go_default_library",
  ],
)

go_binary(
  name = "main_insight",
  srcs = ["main_insight.go"],
  deps = [
    ":insight",
    "//analysis/detectors:sma_crossover",
    "//protos:quote_go_proto",
    "@org_golang_google_grpc//:go_default_library",
  ],
)
