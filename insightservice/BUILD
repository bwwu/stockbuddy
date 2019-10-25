load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "insightservice_lib",
    importpath = "stockbuddy/insightservice",
    srcs = [
      "insightservice.go",
    ],
    deps = [
      "//protos:insight_go_proto",
      "@org_golang_google_grpc//:go_default_library"
    ],
)

go_binary(
  name = "insightservice",
  embed = [":insightservice_lib"],
)
