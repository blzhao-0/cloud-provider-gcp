load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "subtle",
    srcs = ["aliasing.go"],
    importmap = "k8s.io/cloud-provider-gcp/vendor/golang.org/x/crypto/internal/subtle",
    importpath = "golang.org/x/crypto/internal/subtle",
    visibility = ["//vendor/golang.org/x/crypto:__subpackages__"],
)
