.PHONY: default
default: build

all: build test

version := "0.1.0"

build:
	mkdir -p ./.github/bin
	go build -o ./.github/bin


test: build
	go test -short -coverprofile=./.github/bin/cov.out `go list ./... | grep -v vendor/`
	go tool cover -func=./.github/bin/cov.out


