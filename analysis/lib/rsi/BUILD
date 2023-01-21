load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")
package(default_visibility = ["//visibility:public"])

go_library(
  name = "rsi",
  srcs = ["rsi.go"],
  deps = ["//analysis/lib/sma"],
  importpath = "stockbuddy/analysis/lib/rsi",
)

go_test(
  name = "rsi_test",
  srcs = ["rsi_test.go"],
  deps = [":rsi"],
)
