load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "resource",
    srcs = ["resource.pb.go"],
    importmap = "k8s.io/cloud-provider-gcp/vendor/go.opentelemetry.io/proto/otlp/resource/v1",
    importpath = "go.opentelemetry.io/proto/otlp/resource/v1",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/golang/protobuf/proto",
        "//vendor/go.opentelemetry.io/proto/otlp/common/v1:common",
        "//vendor/google.golang.org/protobuf/reflect/protoreflect",
        "//vendor/google.golang.org/protobuf/runtime/protoimpl",
    ],
)
