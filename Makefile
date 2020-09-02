bzl_gazelle:
	bazel run //:gazelle
	bazel run //:gazelle -- update-repos -from_file=go.mod -to_macro=deps.bzl%go_dependencies

bzl_build_project: bzl_gazelle
	bazel build //:project

bzl_run_project: bzl_build_project
	bazel run //:project

bzl_build_apihttp: bzl_gazelle
	bazel build //apihttp

bzl_run_apihttp: bzl_build_project
	bazel run //apihttp
