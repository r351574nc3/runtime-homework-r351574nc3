load("@bazel_gazelle//:def.bzl", "gazelle")

genrule(
    name = "app",
    srcs = [
        "//cmd:support-cases",
    ],
    outs = [
        "app",
    ],
    cmd = "cp ./bazel-out/*/bin/cmd/*/support-cases \"$@\"",
    executable = 1,
    local = 1,
)

gazelle(
    name = "gazelle",
    prefix = "github.com/heroku/runtime-homework-r351574nc3",
)
