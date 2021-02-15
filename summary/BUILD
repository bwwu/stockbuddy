load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_binary(
  name = "volume_summary",
  srcs = [
    "summary.go",
  ],
  deps = [
    "//protos:quote_go_proto",
    "//smtp",
    "@org_golang_google_grpc//:go_default_library",
  ],
)

go_binary(
  name = "daily",
  srcs = ["daily.go"],
  deps = [
    "//analysis/detectors/pricedelta",
    "//analysis/insight",
    "//protos:quote_go_proto",
    "//smtp",
    "@org_golang_google_grpc//:go_default_library",
  ]
)