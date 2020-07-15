.PHONY: generate build clean test lint dep-update

VERSION ?= $(shell git describe --tags --always --dirty --match=v* 2> /dev/null || \
            echo v0)

generate:
	GOFLAGS=-mod=vendor go generate ./...

build:
	\
		CGO_ENABLED=0 \
		gox \
		-osarch="linux/amd64" \
		-mod vendor \
		-ldflags '-extldflags "-static" -X "main.version=${VERSION}"' \
		.

clean:
	rm -f {{ cookiecutter.name }}
	rm -f {{ cookiecutter.name }}_linux_*

test:
	\
		GOFLAGS=-mod=vendor \
		go -v -race -cover test ./...

lint:
	golangci-lint run --enable-all

dep-update:
	go get -u ./...
	go test ./...
	go mod tidy
	go mod vendor
