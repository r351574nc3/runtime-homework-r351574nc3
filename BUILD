load("@bazel_gazelle//:def.bzl", "gazelle")

genrule(
    name = "app",
    srcs = [
        "//cmd:support-cases",
    ],
    outs = [
        "support-cases",
    ],
    cmd = "cp ./bazel-out/*/bin/cmd/*/support-cases \"$@\" && ls -l",
    executable = 1,
    local = 1,
)

gazelle(
    name = "gazelle",
    prefix = "github.com/r351574nc3/runtime-homework-r351574nc3",
)
