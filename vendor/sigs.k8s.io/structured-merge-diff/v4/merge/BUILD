load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "merge",
    srcs = [
        "conflict.go",
        "update.go",
    ],
    importmap = "k8s.io/cloud-provider-gcp/vendor/sigs.k8s.io/structured-merge-diff/v4/merge",
    importpath = "sigs.k8s.io/structured-merge-diff/v4/merge",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/sigs.k8s.io/structured-merge-diff/v4/fieldpath",
        "//vendor/sigs.k8s.io/structured-merge-diff/v4/typed",
    ],
)
