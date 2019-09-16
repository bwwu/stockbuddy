load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_library(
  name = "moving_average",
  srcs = ["moving_average.go"],
  importpath = "stockbuddy/analysis/moving_average",
  deps = [
    "//protos:quote_go_proto",
  ],
)

go_test(
  name = "moving_average_test",
  srcs = ["moving_average_test.go"],
  deps = [
    ":moving_average",
    "//protos:quote_go_proto",
  ],
)

go_binary(
  name = "main",
  srcs = ["main.go"],
  deps = [
    ":moving_average",
    "//protos:quote_go_proto",
    "@org_golang_google_grpc//:go_default_library",
  ],
)
