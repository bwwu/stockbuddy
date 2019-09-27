load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "yahoo_finance_client",
    srcs = ["yahoo_finance_client.go"],
    importpath = "stockbuddy/quote_service/yahoofinanceclient",
)

go_binary(
    name = "quote_service",
    srcs = ["quote_service.go"],
    deps = [
        ":yahoo_finance_client",
        "//protos:quote_go_proto",
        "@org_golang_google_grpc//:go_default_library",
    ],
)
