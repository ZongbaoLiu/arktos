load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["strategy.go"],
    importpath = "k8s.io/kubernetes/pkg/registry/core/datapartition",
    visibility = ["//visibility:public"],
    deps = [
        "//pkg/api/legacyscheme:go_default_library",
        "//pkg/apis/core:go_default_library",
        "//pkg/apis/core/validation:go_default_library",
        "//staging/src/k8s.io/apimachinery/pkg/runtime:go_default_library",
        "//staging/src/k8s.io/apimachinery/pkg/util/validation/field:go_default_library",
        "//staging/src/k8s.io/apiserver/pkg/registry/rest:go_default_library",
        "//staging/src/k8s.io/apiserver/pkg/storage/names:go_default_library",
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
    srcs = [
        ":package-srcs",
        "//pkg/registry/core/datapartition/storage:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
