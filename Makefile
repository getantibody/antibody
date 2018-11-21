SOURCE_FILES?=./...
TEST_PATTERN?=.
TEST_OPTIONS?=

export PATH := ./bin:$(PATH)
export GO111MODULE := on

# Install all the build and lint dependencies
setup:
	curl -sfL https://install.goreleaser.com/github.com/gohugoio/hugo.sh | sh
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh
	go mod download
.PHONY: setup

# Run all the tests
test:
	go test $(TEST_OPTIONS) -failfast -race -coverpkg=./... -covermode=atomic -coverprofile=coverage.txt $(SOURCE_FILES) -run $(TEST_PATTERN) -timeout=2m
.PHONY: test

# Run all the tests and opens the coverage report
cover: test
	go tool cover -html=coverage.txt
.PHONY: cover

# Run all the linters
lint:
	./bin/golangci-lint run --enable-all ./...
	find . -name '*.md' -not -wholename './vendor/*' | xargs prettier -l
.PHONY: lint

# Run all the tests and code checks
ci: build test lint
.PHONY: ci

# Build a beta version
build:
	go build
.PHONY: build

# gofmt and goimports all go files
fmt:
	find . -name '*.go' -not -wholename './vendor/*' | while read -r file; do gofmt -w -s "$$file"; goimports -w "$$file"; done
	find . -name '*.md' -not -wholename './vendor/*' | xargs prettier --write
.PHONY: fmt

# Generate the static documentation
static:
	@hugo --enableGitInfo --source www
.PHONY: static

serve:
	@hugo server --enableGitInfo --watch --source www
.PHONY: serve

favicon:
	wget -O www/static/avatar.png https://avatars2.githubusercontent.com/u/16625397
	convert www/static/avatar.png -define icon:auto-resize=64,48,32,16 www/static/favicon.ico
	convert www/static/avatar.png -resize x120 www/static/apple-touch-icon.png
.PHONY: favicon

.DEFAULT_GOAL := build
