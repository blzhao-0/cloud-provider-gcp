load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "cobra",
    srcs = [
        "args.go",
        "bash_completions.go",
        "bash_completionsV2.go",
        "cobra.go",
        "command.go",
        "command_notwin.go",
        "command_win.go",
        "completions.go",
        "fish_completions.go",
        "powershell_completions.go",
        "shell_completions.go",
        "zsh_completions.go",
    ],
    importmap = "k8s.io/cloud-provider-gcp/vendor/github.com/spf13/cobra",
    importpath = "github.com/spf13/cobra",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/spf13/pflag",
    ] + select({
        "@io_bazel_rules_go//go/platform:windows": [
            "//vendor/github.com/inconshreveable/mousetrap",
        ],
        "//conditions:default": [],
    }),
)
