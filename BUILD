load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_binary(
    name = "main",
    embed = [":go_default_library"],
    visibility = ["//visibility:public"],
)

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    importpath = "stockbuddy",
    visibility = ["//visibility:private"],
    deps = [
        "@org_mongodb_go_mongo_driver//mongo:go_default_library",
        "@org_mongodb_go_mongo_driver//mongo/options:go_default_library",
    ],
)
