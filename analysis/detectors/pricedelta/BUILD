# pricedelta is a summary indicator, therefore always emits.
load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

package(default_visibility = ["//visibility:public"])

go_library(
  name = "pricedelta",
  srcs = ["pricedelta.go"],
  importpath = "stockbuddy/analysis/detectors/pricedelta",
  deps = [
    "//analysis/constants",
    "//analysis/insight",
    "//protos:quote_go_proto",
  ],
)