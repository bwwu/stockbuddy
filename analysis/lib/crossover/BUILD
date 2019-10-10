load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

package(default_visibility = ["//visibility:public"])

go_library(
    name = "crossover",
    srcs = ["crossover.go"],
    importpath = "stockbuddy/analysis/lib/crossover",
    deps = ["//analysis:constants"],
)

go_test(
    name = "crossover_test",
    srcs = ["crossover_test.go"],
    deps = [
        ":crossover",
        "//analysis:constants",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["crossover_test.go"],
    embed = [":go_default_library"],
    deps = [
        ":crossover",
        "//analysis:constants",
    ],
)
