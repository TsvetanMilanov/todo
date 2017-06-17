SHELL := bash

rwildcard = $(wildcard $1$2) $(foreach d,$(wildcard $1*),$(call rwildcard,$d/,$2))
executablename = $(shell [[ $(OS) = 'Windows_NT' ]] && echo $1.exe || echo $1)

BUILDDIR  = .build
CLIENTSDIR = clients
SERVERDIR = server
CLIDIR = $(CLIENTSDIR)/cli

.PHONY: \
	build-cli \
	glide-all \
	build-server \
	clean/%

$(BUILDDIR)/glide-install-cli: $(CLIDIR)/glide.yaml
	@cd $(CLIDIR) && glide install
	@touch $@

$(BUILDDIR)/glide-install-server: $(SERVERDIR)/glide.yaml
	@cd $(SERVERDIR) && glide install
	@touch $@

glide-all: $(BUILDDIR)/glide-install-cli $(BUILDDIR)/glide-install-server
	@echo Installed all dependencies.

$(BUILDDIR)/build-cli: $(BUILDDIR)/glide-install-cli $(call rwildcard,$(CLIDIR)/cmd,*.go) $(call rwildcard,$(CLIDIR)/pkg,*.go) $(CLIDIR)/main.go
	@go build -o $(GOPATH)/bin/$(call executablename,todo) $(CLIDIR)/main.go
	@touch $@

build-cli: $(BUILDDIR)/build-cli
	@echo TODO CLI successfully built.

$(BUILDDIR)/build-server: $(BUILDDIR)/glide-install-server $(call rwildcard,$(SERVERDIR)/pkg,*.go) $(SERVERDIR)/main.go
	@go build -o $(SERVERDIR)/$(call executablename,main) $(SERVERDIR)/main.go
	@touch $@

build-server: $(BUILDDIR)/build-server
	@echo TODO server successfully built.

clean/%:
	rm -rf $(BUILDDIR)/*$**
