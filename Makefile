SHELL := bash
rwildcard = $(wildcard $1$2) $(foreach d,$(wildcard $1*),$(call rwildcard,$d/,$2))
BUILDDIR  = .build
CLIENTSDIR = clients
SERVERDIR = server
CLIDIR = $(CLIENTSDIR)/cli

.PHONY: \
	build-cli \
	glide-all

$(BUILDDIR)/glide-install-cli: $(CLIDIR)/glide.yaml
	@cd $(CLIDIR) && glide install
	@touch $@

$(BUILDDIR)/glide-install-server: $(SERVERDIR)/glide.yaml
	@cd $(SERVERDIR) && glide install
	@touch $@

glide-all: $(BUILDDIR)/glide-install-cli $(BUILDDIR)/glide-install-server
	@echo Installed all dependencies.

$(BUILDDIR)/build-cli: $(BUILDDIR)/glide-install-cli $(call rwildcard,$(CLIDIR)/cmd,*.go) $(call rwildcard,$(CLIDIR)/pkg,*.go) $(CLIDIR)/main.go
	@go build -o $(GOPATH)/bin/todo $(CLIDIR)/main.go
	@touch $@

build-cli: $(BUILDDIR)/build-cli
	@echo TODO CLI successfully built.
