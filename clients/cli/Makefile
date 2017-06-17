SHELL := bash
rwildcard = $(wildcard $1$2) $(foreach d,$(wildcard $1*),$(call rwildcard,$d/,$2))
BUILDDIR  = .build

.PHONY: \
	build-cli

$(BUILDDIR)/build-cli: $(call rwildcard,cmd,*.go) $(call rwildcard,pkg,*.go) main.go
	@go build -o $(GOPATH)/bin/todo ./
	@touch $@

build-cli: $(BUILDDIR)/build-cli
	@echo TODO CLI successfully built.
