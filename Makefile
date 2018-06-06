SOURCE_FILES?=./...
TEST_PATTERN?=.
TEST_OPTIONS?=
OS=$(shell uname -s)

export PATH := ./bin:$(PATH)

# Install all the build and lint dependencies
setup:
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh
ifeq ($(OS), Darwin)
	brew install dep
else
	curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
endif
	dep ensure -vendor-only
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
	golangci-lint run --enable-all ./...
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

# Generates the static documentation
static-gen:
	@rm -rf dist/getantibody.github.io/theme
	@static-docs \
		--in docs \
		--out dist/getantibody.github.io \
		--title Antibody \
		--subtitle "The fastest shell plugin manager" \
		--google UA-68164063-1
.PHONY: static-gen

# Downloads and generates the static documentation
static:
	@rm -rf dist/getantibody.github.io
	@mkdir -p dist
	@git clone https://github.com/getantibody/getantibody.github.io.git dist/getantibody.github.io
	@make static-gen
.PHONY: static

.DEFAULT_GOAL := build
