load("@io_bazel_rules_go//go:def.bzl", "go_library")
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")

proto_library(
    name = "quote_proto",
    srcs = ["quote.proto"],
    visibility = ["//visibility:public"],
)

go_proto_library(
    name = "quote_go_proto_lib",
    compilers = ["@io_bazel_rules_go//proto:go_grpc"],
    importpath = "stockbuddy/protos/quote",
    proto = ":quote_proto",
    visibility = ["//visibility:public"],
)

go_library(
    name = "quote_go_proto",
    embed = [":quote_go_proto"],
    importpath = "stockbuddy/protos/quote",
    visibility = ["//visibility:public"],
)
