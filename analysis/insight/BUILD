load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
  name = "insight",
  srcs = [
    "insight.go",
    "insight_formatter.go",
  ],
  importpath = "stockbuddy/analysis/insight",
  visibility = ["//visibility:public"],
  deps = [
    "//analysis/constants",
    "//protos:quote_go_proto",
  ],
)


