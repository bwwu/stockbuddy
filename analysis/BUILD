load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")


go_library(
  name = "stocks_to_watch",
  importpath = "stockbuddy/analysis/stocks_to_watch",
  srcs = ["stocks_to_watch.go"],
)

go_binary(
  name = "main",
  srcs = ["main.go"],
  deps = [
    ":stocks_to_watch",
    "//analysis/moving_average:moving_average",
    "//analysis/moving_average/crossover:crossover",
    "//analysis/moving_average/crossover:crossover_reporter",
    "//protos:quote_go_proto",
    "//smtp:sendmail",
    "@org_golang_google_grpc//:go_default_library",
  ],
)
