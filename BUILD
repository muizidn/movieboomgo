load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix github.com/muizidn/movieboomgo
gazelle(name = "gazelle")

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    importpath = "github.com/muizidn/movieboomgo",
    visibility = ["//visibility:private"],
)

go_binary(
    name = "project",
    embed = [":go_default_library"],
    visibility = ["//visibility:public"],
)
