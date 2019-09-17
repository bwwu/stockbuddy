load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_binary(
  name = "main",
  srcs = ["main.go"],
  deps = [
    "//analysis/moving_average:moving_average",
    "//protos:quote_go_proto",
    "//smtp:sendmail",
    "@org_golang_google_grpc//:go_default_library",
  ],
)
