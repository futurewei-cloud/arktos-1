load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["event-system.go"],
    importmap = "k8s.io/kubernetes/vendor/github.com/go-chassis/go-archaius/core/event-system",
    importpath = "github.com/go-chassis/go-archaius/core/event-system",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/go-chassis/go-archaius/core:go_default_library",
        "//vendor/github.com/go-mesh/openlogging:go_default_library",
    ],
)

filegroup(
    name = "package-srcs",
    srcs = glob(["**"]),
    tags = ["automanaged"],
    visibility = ["//visibility:private"],
)

filegroup(
    name = "all-srcs",
    srcs = [":package-srcs"],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
