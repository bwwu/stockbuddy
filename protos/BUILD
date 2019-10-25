load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")

package(default_visibility = ["//visibility:public"])

proto_library(
  name = "quote_proto",
  srcs = ["quote.proto"],
)

go_proto_library(
  name = "quote_go_proto",
  compiler = "@io_bazel_rules_go//proto:go_grpc",
  importpath ="stockbuddy/protos/quote_go_proto",
  proto = ":quote_proto",
)

proto_library(
  name = "insight_proto",
  srcs = ["insight.proto"],
)

go_proto_library(
  name = "insight_go_proto",
  compiler = "@io_bazel_rules_go//proto:go_grpc",
  importpath ="stockbuddy/protos/insight_go_proto",
  proto = ":insight_proto",
)
