load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_go",
    urls = ["https://github.com/bazelbuild/rules_go/releases/download/0.12.1/rules_go-0.12.1.tar.gz"],
    sha256 = "8b68d0630d63d95dacc0016c3bb4b76154fe34fca93efd65d1c366de3fcb4294",
)

http_archive(
    name = "bazel_gazelle",
    urls = ["https://github.com/bazelbuild/bazel-gazelle/releases/download/0.12.0/bazel-gazelle-0.12.0.tar.gz"],
    sha256 = "ddedc7aaeb61f2654d7d7d4fd7940052ea992ccdb031b8f9797ed143ac7e8d43",
)

load("@io_bazel_rules_go//go:def.bzl", "go_rules_dependencies", "go_register_toolchains")

go_rules_dependencies()

go_register_toolchains()

load("@bazel_gazelle//:deps.bzl", "go_repository", "gazelle_dependencies")

gazelle_dependencies()

go_repository(
    name = "com_github_gin_gonic_gin",
    commit = "caf3e350a548af1add9def68087ac53d1d000caa",
    importpath = "github.com/gin-gonic/gin",
)

go_repository(
    name = "com_github_gin_contrib_sse",
    commit = "22d885f9ecc78bf4ee5d72b937e4bbcdc58e8cae",
    importpath = "github.com/gin-contrib/sse",
)

go_repository(
    name = "com_github_mattn_go_isatty",
    commit = "6ca4dbf54d38eea1a992b3c722a76a5d1c4cb25c",
    importpath = "github.com/mattn/go-isatty",
)

go_repository(
    name = "com_github_go_yaml_yaml",
    commit = "5420a8b6744d3b0345ab293f6fcba19c978f1183",
    importpath = "github.com/go-yaml/yaml",
)

go_repository(
    name = "in_gopkg_yaml_v2",
    commit = "5420a8b6744d3b0345ab293f6fcba19c978f1183",
    importpath = "gopkg.in/yaml.v2",
)

go_repository(
    name = "com_github_ugorji_go",
    commit = "f3cacc17c85ecb7f1b6a9e373ee85d1480919868",
    importpath = "github.com/ugorji/go",
)

go_repository(
    name = "in_gopkg_go_playground_validator_v8",
    commit = "5f1438d3fca68893a817e4a66806cea46a9e4ebf",
    importpath = "gopkg.in/go-playground/validator.v8",
)

go_repository(
    name = "com_github_go_playground_validator",
    commit = "f28cb45dc0f0f406ba2015cdb59a5203ad220bc2",
    importpath = "github.com/go-playground/validator",
)

go_repository(
    name = "com_github_go_playground_universal_translator",
    commit = "71201497bace774495daed26a3874fd339e0b538",
    importpath = "github.com/go-playground/universal-translator",
)

go_repository(
    name = "com_github_go_playground_locales",
    commit = "f63010822830b6fe52288ee52d5a1151088ce039",
    importpath = "github.com/go-playground/locales",
)

go_repository(
    name = "com_github_golang_protobuf",
    commit = "5831880292e721c76b58a16ecc60adc27d8e6355",
    importpath = "github.com/golang/protobuf/proto",
)
