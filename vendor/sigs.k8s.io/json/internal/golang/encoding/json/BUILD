load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "json",
    srcs = [
        "decode.go",
        "encode.go",
        "fold.go",
        "fuzz.go",
        "indent.go",
        "kubernetes_patch.go",
        "scanner.go",
        "stream.go",
        "tables.go",
        "tags.go",
    ],
    importmap = "k8s.io/cloud-provider-gcp/vendor/sigs.k8s.io/json/internal/golang/encoding/json",
    importpath = "sigs.k8s.io/json/internal/golang/encoding/json",
    visibility = ["//vendor/sigs.k8s.io/json:__subpackages__"],
)
