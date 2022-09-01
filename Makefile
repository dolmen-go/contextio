

go ?= GO111MODULE=on go
export go

go-version: go.mod $(shell $(go) list -f '{{$$Dir := .Dir}}{{range .GoFiles}}{{$$Dir}}/{{.}} {{end}}' ./...)
	@TZ=UTC git log -1 '--date=format-local:%Y%m%d%H%M%S' --abbrev=12 '--pretty=tformat:v0.0.0-%cd-%h' $^

go-get:
	@echo $(go) get $(shell $(go) list .)@$(shell $(MAKE) -f $(firstword $(MAKEFILE_LIST)) go-version)

