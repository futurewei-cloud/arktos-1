load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["fake_cache.go"],
    importpath = "k8s.io/kubernetes/resourcecollector/pkg/collector/internal/cache/fake",
    visibility = ["//resourcecollector/pkg/collector:__subpackages__"],
    deps = [
        "//globalscheduler/pkg/scheduler/client/typed:go_default_library",
        "//globalscheduler/pkg/scheduler/listers:go_default_library",
        "//globalscheduler/pkg/scheduler/types:go_default_library",
        "//resourcecollector/pkg/collector/internal/cache:go_default_library",
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
