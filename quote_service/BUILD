load("@io_bazel_rules_docker//container:container.bzl", "container_image")
load("@io_bazel_rules_docker//go:image.bzl", "go_image", "DEFAULT_BASE")
load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "yahoo_finance_client",
    srcs = ["yahoo_finance_client.go"],
    importpath = "stockbuddy/quote_service/yahoofinanceclient",
)

container_image(
    name = "base_image",
    base = DEFAULT_BASE,
    ports = ["50051"],
)

go_image(
    name = "image",
    srcs = ["quote_service.go"],
    base = ":base_image",
    deps = [
        "//protos:quote_go_proto",
        "//quote_service/yahoofinance:yahoo_finance_client",
        "@org_golang_google_grpc//:go_default_library",
    ],
)

go_binary(
    name = "quote_service",
    srcs = ["quote_service.go"],
    deps = [
        "//protos:quote_go_proto",
        "//quote_service/yahoofinance:yahoo_finance_client",
        "@org_golang_google_grpc//:go_default_library",
    ],
)
