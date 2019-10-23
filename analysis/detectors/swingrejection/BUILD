load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

package(default_visibility = ["//visibility:public"])

go_library(
  name = "swing_rejection",
  srcs = ["swing_rejection.go"],
  importpath = "stockbuddy/analysis/detectors/swingrejection",
  deps = [
    "//analysis/constants",
    "//analysis/insight",
    "//analysis/lib/rsi",
    "//protos:quote_go_proto",
  ],
)
