load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "fake",
    srcs = [
        "clientset_generated.go",
        "doc.go",
        "register.go",
    ],
    importpath = "k8s.io/cloud-provider-gcp/crd/client/network/clientset/versioned/fake",
    visibility = ["//visibility:public"],
    deps = [
        "//crd/apis/network/v1",
        "//crd/client/network/clientset/versioned",
        "//crd/client/network/clientset/versioned/typed/network/v1:network",
        "//crd/client/network/clientset/versioned/typed/network/v1/fake",
        "//vendor/k8s.io/apimachinery/pkg/apis/meta/v1:meta",
        "//vendor/k8s.io/apimachinery/pkg/runtime",
        "//vendor/k8s.io/apimachinery/pkg/runtime/schema",
        "//vendor/k8s.io/apimachinery/pkg/runtime/serializer",
        "//vendor/k8s.io/apimachinery/pkg/util/runtime",
        "//vendor/k8s.io/apimachinery/pkg/watch",
        "//vendor/k8s.io/client-go/discovery",
        "//vendor/k8s.io/client-go/discovery/fake",
        "//vendor/k8s.io/client-go/testing",
    ],
)
