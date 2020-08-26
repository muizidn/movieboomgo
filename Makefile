bzl_gazelle:
	bazel run //:gazelle

bzl_build_project:
	bazel build //:project

bzl_run_project: bzl_build_project
	bazel-bin/project_/project	
