

go ?= GO111MODULE=on go
export go

tag_prefix :=
go_files = go.mod $(shell $(go) list -f '{{$$Dir := .Dir}}{{range .GoFiles}}{{$$Dir}}/{{.}} {{end}}' ./...)
go_files_last_commit = $(shell git log -1 --format=%H -- $(go_files))

## Show module version as expected by the Go toolchain
go-version: $(go_files)
	@{ git describe --tags --match '$(tag_prefix)v*.*.*' --exact-match 2>/dev/null $(shell git log -1 --format=%H -- $^ ) || TZ=UTC git log -1 '--date=format-local:%Y%m%d%H%M%S' --abbrev=12 '--pretty=tformat:%(describe:tags,match=$(tag_prefix)v*,abbrev=0)-0-%cd-%h' $^ | perl -pE 's/(\d+)(?=-)/$$1+1/e' ; } | sed -e 's!.*/!!'

go-get:
	@echo $(go) get $(shell $(go) list .)@$(shell $(MAKE) -f $(firstword $(MAKEFILE_LIST)) go-version)

