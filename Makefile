.DEFAULT_GOAL = build
GO := go
PKG_NAME := gobot
DOCKER_REPO ?= awayfromserver/$(PKG_NAME)
PREFIX := .
DOCKER_LINUX_PLATFORMS ?= linux/amd64
DOCKER_PLATFORMS ?= $(DOCKER_LINUX_PLATFORMS)
BUILDX_ACTION ?= --load # Default dry run
TAG_LATEST ?= latest
TAG_ALPINE ?= alpine

ifeq ("$(CI)","true")
LINT_PROCS ?= 1
else
LINT_PROCS ?= $(shell nproc)
endif

COMMIT ?= `git rev-parse --short HEAD 2>/dev/null`
VERSION ?= $(shell $(GO) run ./version/gen/vgen.go)
VERSION_PATH ?= `$(GO) list ./version`
COMMIT_FLAG ?= -X $(VERSION_PATH).GitCommit=$(COMMIT)
VERSION_FLAG ?= -X $(VERSION_PATH).Version=$(VERSION)
GO_LDFLAGS ?= $(COMMIT_FLAG) $(VERSION_FLAG)

GOOS ?= $(shell $(GO) version | sed 's/^.*\ \([a-z0-9]*\)\/\([a-z0-9]*\)/\1/')
GOARCH ?= $(shell $(GO) version | sed 's/^.*\ \([a-z0-9]*\)\/\([a-z0-9]*\)/\2/')
CGO_ENABLED=0


clean:
	rm -Rf $(PREFIX)/bin/*
	rm -f $(PREFIX)/*.[ci]id

build: $(PREFIX)/bin/$(PKG_NAME)_$(GOOS)-$(GOARCH)$(TARGETVARIANT)$(call extension,$(GOOS)) $(PREFIX)/bin/$(PKG_NAME)$(call extension,$(GOOS))

test:
	$(GO) test -race -coverprofile=c.out ./...

lint:
	@golangci-lint run --verbose --max-same-issues=0 --max-issues-per-linter=0

ci-lint:
	@golangci-lint run --verbose --max-same-issues=0 --max-issues-per-linter=0 --out-format=github-actions

.PHONY: clean test build lint clean-images clean-containers docker-images
.DELETE_ON_ERROR: