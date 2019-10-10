load("@io_bazel_rules_go//go:def.bzl", "go_library")
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")

package(default_visibility = ["//visibility:public"])

proto_library(
    name = "quote_proto",
    srcs = ["quote.proto"],
)

go_proto_library(
    name = "quote_go_proto",
    compiler = "@io_bazel_rules_go//proto:go_grpc",
    compilers = ["@io_bazel_rules_go//proto:go_grpc"],
    importpath = "stockbuddy/protos/quote_go_proto",
    proto = ":quote_go_proto_proto",
)

proto_library(
    name = "quote_go_proto_proto",
    srcs = ["quote.proto"],
)

go_library(
    name = "go_default_library",
    embed = [":quote_go_proto"],
    importpath = "stockbuddy/protos/quote_go_proto",
)
