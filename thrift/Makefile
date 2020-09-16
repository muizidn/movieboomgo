gen_thrift_go:
	thrift -r --gen go:package_prefix=github.com/muizidn/movieboomgo/thrift/gen-go/ movie.thrift

gen_thrift_swift:
	thrift -r --gen swift:safe_enums movie.thrift

all: gen_thrift_go gen_thrift_swift

.PHONY: all
	