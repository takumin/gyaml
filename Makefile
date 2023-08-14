APPNAME  := $(shell basename $(CURDIR))
PKGNAME  := $(shell go mod edit -json | jq -r '.Module.Path')
VERSION  := $(shell git describe --abbrev=0 --tags 2>/dev/null)
REVISION := $(shell git rev-parse HEAD 2>/dev/null)

ifeq ($(VERSION),)
VERSION := dev
endif

ifeq ($(REVISION),)
REVISION := unknown
endif

SRCS := $(shell find . -type f -name '*.go')

LDFLAGS_APPNAME  := -X "$(PKGNAME)/internal/metadata.appName=$(APPNAME)"
LDFLAGS_VERSION  := -X "$(PKGNAME)/internal/version.version=$(VERSION)"
LDFLAGS_REVISION := -X "$(PKGNAME)/internal/version.revision=$(REVISION)"
LDFLAGS          := -s -w -buildid= $(LDFLAGS_APPNAME) $(LDFLAGS_VERSION) $(LDFLAGS_REVISION) -extldflags -static
BUILDFLAGS       := -trimpath -ldflags '$(LDFLAGS)'

.PHONY: all
all: clean tools generate fmt vet sec vuln lint test build

.PHONY: docs
docs:
	mdbook build

.PHONY: docs-serve
docs-serve:
	mdbook serve -o

.PHONY: tools
tools:
	aqua install --all

.PHONY: generate
generate:
	go generate ./...

.PHONY: fmt
fmt:
	gofmt -e -d .

.PHONY: vet
vet:
	go vet ./...

.PHONY: sec
sec:
	gosec -quiet -fmt golint ./...

.PHONY: vuln
vuln:
	trivy fs -q -s HIGH,CRITICAL --scanners vuln,config,secret,license .

.PHONY: lint
lint:
	staticcheck ./...

.PHONY: test
test:
	go test -trimpath -cover -covermode atomic ./...

.PHONY: coverage
coverage: coverage.out coverage.html
coverage.out: $(SRCS)
	go test -trimpath -cover -covermode atomic -coverprofile=coverage.out ./...
coverage.html: coverage.out
	go tool cover -html=coverage.out -o coverage.html

.PHONY: build
build: bin/$(APPNAME)
bin/$(APPNAME): $(SRCS)
	CGO_ENABLED=0 go build $(BUILDFLAGS) -o $@

.PHONY: install
install: build
	CGO_ENABLED=0 go install $(BUILDFLAGS)

.PHONY: snapshot
snapshot:
	goreleaser release --rm-dist --snapshot

.PHONY: release
release:
ifneq ($(GITHUB_TOKEN),)
	goreleaser release --rm-dist
endif

.PHONY: clean
clean:
	rm -rf bin
	rm -rf dist
	rm -rf book
	rm -f coverage.out
	rm -f coverage.html
