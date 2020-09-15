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

bzl_build_thrift: bzl_gazelle
	bazel build //thrift

bzl_run_thrift_client: bzl_build_thrift
	./bazel-bin/thrift/thrift_/thrift

bzl_run_thrift_server: bzl_build_thrift
	./bazel-bin/thrift/thrift_/thrift -server