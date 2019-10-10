load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "stocks_to_watch",
    srcs = ["stocks_to_watch.go"],
    importpath = "stockbuddy/analysis/stocks_to_watch",
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
        "insight_formatter.go",
    ],
    importpath = "stockbuddy/analysis/insight",
    visibility = ["//visibility:public"],
    deps = [
        "//analysis:constants",
        "//protos:quote_go_proto",
    ],
)

go_binary(
    name = "analysis",
    srcs = [
        "main.go",
        "stocks_to_watch.go",
    ],
    deps = [
        ":insight",
        ":stocks_to_watch",
        "//analysis/detectors:macd_crossover",
        "//analysis/detectors:sma_crossover",
        "//analysis/detectors:swing_rejection",
        "//protos:quote_go_proto",
        "//smtp",
        "@org_golang_google_grpc//:go_default_library",
    ],
)
