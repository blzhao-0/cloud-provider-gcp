load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "util",
    srcs = [
        "build.go",
        "plural_exceptions.go",
    ],
    importmap = "k8s.io/cloud-provider-gcp/vendor/k8s.io/code-generator/pkg/util",
    importpath = "k8s.io/code-generator/pkg/util",
    visibility = ["//visibility:public"],
    deps = ["//vendor/golang.org/x/tools/go/packages"],
)
