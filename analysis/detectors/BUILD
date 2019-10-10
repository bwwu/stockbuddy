load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

package(default_visibility = ["//visibility:public"])

go_library(
    name = "macd_crossover",
    srcs = ["macd_crossover.go"],
    importpath = "stockbuddy/analysis/detectors/macd_crossover",
    deps = [
        "//analysis:constants",
        "//analysis:insight",
        "//analysis/lib:macd",
        "//analysis/lib:sma",
        "//analysis/lib/crossover",
        "//protos:quote_go_proto",
    ],
)

go_test(
    name = "macd_crossover_test",
    srcs = ["macd_crossover_test.go"],
    deps = [
        ":macd_crossover",
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
        "//analysis/lib/crossover",
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

go_library(
    name = "swing_rejection",
    srcs = ["swing_rejection.go"],
    importpath = "stockbuddy/analysis/detectors/swing_rejection",
    deps = [
        "//analysis:constants",
        "//analysis:insight",
        "//analysis/lib:rsi",
        "//protos:quote_go_proto",
    ],
)

go_library(
    name = "go_default_library",
    srcs = [
        "macd_crossover.go",
        "sma_crossover.go",
        "swing_rejection.go",
    ],
    importpath = "stockbuddy/analysis/detectors",
    deps = [
        "//analysis:constants",
        "//analysis:insight",
        "//analysis/lib:macd",
        "//analysis/lib:rsi",
        "//analysis/lib:sma",
        "//analysis/lib/crossover",
        "//protos:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["sma_crossover_test.go"],
    embed = [":go_default_library"],
    deps = [
        ":sma_crossover",
        "//analysis:constants",
        "//protos:go_default_library",
    ],
)
