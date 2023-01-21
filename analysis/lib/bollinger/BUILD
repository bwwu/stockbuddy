load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

package(default_visibility = ["//visibility:public"])

go_library(
  name = "bollinger",
  srcs = ["bollinger.go"],
  importpath = "stockbuddy/analysis/lib/bollinger",
  deps = [
    "//analysis/lib/sma",
  ],
)

go_test(
  name = "bollinger_test",
  srcs = ["bollinger_test.go"],
  deps = [":bollinger"],
)
