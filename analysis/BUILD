load("@io_bazel_rules_go//go:def.bzl", "go_binary")

go_binary(
  name = "movingaverage",
  srcs = ["movingaverage.go"],
  deps = [
    "//protos:quote_go_proto",
    "@org_golang_google_grpc//:go_default_library",
  ],
)
