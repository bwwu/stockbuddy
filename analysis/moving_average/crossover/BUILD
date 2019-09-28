load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

package(default_visibility = ["//visibility:public"])

go_library(
  name = "crossover",
  srcs = ["crossover.go"],
  importpath = "stockbuddy/analysis/moving_average/crossover/crossover",
  deps = ["//analysis:constants"],
)

go_library(
  name = "crossover_reporter",
  srcs = ["crossover_reporter.go"],
  importpath = "stockbuddy/analysis/moving_average/crossover/crossover_reporter",
  deps = [
    ":crossover",
    "//analysis:constants",
  ],
)

go_test(
  name = "crossover_test",
  srcs = ["crossover_test.go"],
  deps = [
    ":crossover",
    "//analysis:constants",
  ],
)
