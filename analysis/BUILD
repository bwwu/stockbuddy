load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_binary(
  name = "analysis",
  srcs = [
    "main.go",
    "stocks_to_watch.go",
  ],
  deps = [
    "//analysis/insight",
    "//analysis/detectors/smax:sma_crossover",
    "//analysis/detectors/macdx:macd_crossover",
    "//analysis/detectors/swingrejection:swing_rejection",
    "//protos:quote_go_proto",
    "//smtp",
    "@org_golang_google_grpc//:go_default_library",
  ],
)
