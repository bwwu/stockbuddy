load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

package(default_visibility = ["//visibility:public"])

go_library(
  name = "percentile",
  srcs = ["percentile.go"],
  importpath = "stockbuddy/analysis/lib/percentile",
)

go_test(
  name = "percentile_test",
  srcs = ["percentile_test.go"],
  deps = [
    ":percentile",
  ],
)
