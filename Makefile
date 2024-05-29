.DEFAULT_GOAL = build
GO := go
extension = $(patsubst windows,.exe,$(filter windows,$(1)))
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

ifeq ("$(TARGETVARIANT)","")
ifneq ("$(GOARM)","")
TARGETVARIANT := v$(GOARM)
endif
else
ifeq ("$(GOARM)","")
GOARM ?= $(subst v,,$(TARGETVARIANT))
endif
endif

platforms := linux-amd64

clean:
	rm -Rf $(PREFIX)/bin/*
	rm -f $(PREFIX)/*.[ci]id

build-x: $(patsubst %,$(PREFIX)/bin/$(PKG_NAME)_%,$(platforms))

$(PREFIX)/bin/%.zip: $(PREFIX)/bin/%
	@zip -j $@ $^

$(PREFIX)/bin/$(PKG_NAME)_%_checksum_sha256.txt: $(PREFIX)/bin/$(PKG_NAME)_%
	@sha256sum $< > $@

$(PREFIX)/bin/$(PKG_NAME)_%_checksum_sha512.txt: $(PREFIX)/bin/$(PKG_NAME)_%
	@sha512sum $< > $@

$(PREFIX)/bin/checksums.txt: $(PREFIX)/bin/checksums_sha256.txt
	@cp $< $@

$(PREFIX)/bin/checksums_sha256.txt: \
		$(patsubst %,$(PREFIX)/bin/$(PKG_NAME)_%_checksum_sha256.txt,$(platforms))
	@cat $^ > $@

$(PREFIX)/bin/checksums_sha512.txt: \
		$(patsubst %,$(PREFIX)/bin/$(PKG_NAME)_%_checksum_sha512.txt,$(platforms))
	@cat $^ > $@

$(PREFIX)/%.signed: $(PREFIX)/%
	@keybase sign < $< > $@

%.iid: Dockerfile
	@docker build \
		--build-arg VCS_REF=$(COMMIT) \
		--target $(subst .iid,,$@) \
		--iidfile $@ \
		.

docker-multi: Dockerfile
	docker buildx build \
		--build-arg VCS_REF=$(COMMIT) \
		--platform $(DOCKER_PLATFORMS) \
		--tag $(DOCKER_REPO):$(TAG_LATEST) \
		--target gobot \
		$(BUILDX_ACTION) .
	docker buildx build \
		--build-arg VCS_REF=$(COMMIT) \
		--platform $(DOCKER_LINUX_PLATFORMS) \
		--tag $(DOCKER_REPO):$(TAG_ALPINE) \
		--target gobot-alpine \
		$(BUILDX_ACTION) .

%.cid: %.iid
	@docker create $(shell cat $<) > $@

build-release: artifacts.cid
	@docker cp $(shell cat $<):/bin/. bin/

docker-images: gobot.iid

GO_FILES := $(shell find . -type f -name "*.go")

$(PREFIX)/bin/$(PKG_NAME)_%$(TARGETVARIANT)$(call extension,$(GOOS)): $(GO_FILES)
	GOOS=$(shell echo $* | cut -f1 -d-) GOARCH=$(shell echo $* | cut -f2 -d- ) GOARM=$(GOARM) CGO_ENABLED=$(CGO_ENABLED) \
		$(GO) build \
			-ldflags "-w -s $(GO_LDFLAGS)" \
			-o $@ \
			.

$(PREFIX)/bin/$(PKG_NAME)$(call extension,$(GOOS)): $(PREFIX)/bin/$(PKG_NAME)_$(GOOS)-$(GOARCH)$(TARGETVARIANT)$(call extension,$(GOOS))
	cp $< $@

build: $(PREFIX)/bin/$(PKG_NAME)_$(GOOS)-$(GOARCH)$(TARGETVARIANT)$(call extension,$(GOOS)) $(PREFIX)/bin/$(PKG_NAME)$(call extension,$(GOOS))

test:
	$(GO) test -race -coverprofile=c.out ./...

integration: $(PREFIX)/bin/$(PKG_NAME)
	$(GO) test -v \
		-ldflags "-X `$(GO) list ./internal/tests/integration`.GobotBinPath=$(shell pwd)/$<" \
		./internal/tests/integration

integration.iid: Dockerfile.integration $(PREFIX)/bin/$(PKG_NAME)_linux-amd64$(call extension,$(GOOS))
	docker build -f $< --iidfile $@ .

test-integration-docker: integration.iid
	docker run -it --rm $(shell cat $<)

lint:
	@golangci-lint run --verbose --max-same-issues=0 --max-issues-per-linter=0

ci-lint:
	@golangci-lint run --verbose --max-same-issues=0 --max-issues-per-linter=0 --out-format=github-actions

.PHONY: clean test build-x build-release build test-integration-docker lint clean-images clean-containers docker-images integration
.DELETE_ON_ERROR:
