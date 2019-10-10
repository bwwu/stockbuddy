load("@io_bazel_rules_go//go:def.bzl", "go_library")
load("@io_bazel_rules_go//proto:def.bzl", "go_grpc_library", "go_proto_library")

package(default_visibility = ["//visibility:public"])

proto_library(
    name = "quote_proto",
    srcs = ["quote.proto"],
)

go_proto_library(
    name = "quote_go_proto",
    compilers = ["@io_bazel_rules_go//proto:go_grpc"],
    importpath = "stockbuddy/protos/quote_go_proto",
    proto = ":quote_proto",
)
