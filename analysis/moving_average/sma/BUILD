load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

package(default_visibility = ["//visibility:public"])

go_library(
  name = "sma",
  srcs = ["sma.go"],
  importpath = "stockbuddy/analysis/moving_average/sma/sma",
)

go_test(
  name = "sma_test",
  srcs = ["sma_test.go"],
  deps = [":sma"],
)
