load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "gcp",
    srcs = ["gcp.go"],
    importmap = "k8s.io/cloud-provider-gcp/vendor/k8s.io/client-go/plugin/pkg/client/auth/gcp",
    importpath = "k8s.io/client-go/plugin/pkg/client/auth/gcp",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/golang.org/x/oauth2",
        "//vendor/golang.org/x/oauth2/google",
        "//vendor/k8s.io/apimachinery/pkg/util/net",
        "//vendor/k8s.io/apimachinery/pkg/util/yaml",
        "//vendor/k8s.io/client-go/rest",
        "//vendor/k8s.io/client-go/util/jsonpath",
        "//vendor/k8s.io/klog/v2:klog",
    ],
)
