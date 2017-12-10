SOURCE_FILES?=./...
TEST_PATTERN?=.
TEST_OPTIONS?=

# Install all the build and lint dependencies
setup:
	go get -u github.com/alecthomas/gometalinter
	go get -u github.com/golang/dep/cmd/dep
	go get -u github.com/pierrre/gotestcover
	go get -u golang.org/x/tools/cmd/cover
	go get -u github.com/apex/static/cmd/static-docs
	dep ensure
	gometalinter --install
.PHONY: setup

# Run all the tests
test:
	gotestcover $(TEST_OPTIONS) -covermode=atomic -coverprofile=coverage.txt $(SOURCE_FILES) -run $(TEST_PATTERN) -timeout=60s
.PHONY: test

# Run all the tests and opens the coverage report
cover: test
	go tool cover -html=coverage.txt
.PHONY: cover

# Run all the linters
lint:
	gometalinter --vendor ./...
.PHONY: lint

# Run all the tests and code checks
ci: lint test
.PHONY: ci

# Build a beta version
build:
	go build
.PHONY: build

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

# Opens the current docs on the default browser
static-open:
	open dist/getantibody.github.io/index.html
.PHONY: static-open

.DEFAULT_GOAL := build
