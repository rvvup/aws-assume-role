REVISION := $(shell git rev-parse HEAD)
CHANGES := $(shell test -n "$$(git status --porcelain)" && echo '+CHANGES' || true)

LDFLAGS := -X main.Revision=$(REVISION)$(CHANGES) -X main.version="0.1.0"

build: build-linux build-osx build-osx-arm

build-linux:
	@ GOOS=linux CGO_ENABLED=0 go build -ldflags="$(LDFLAGS)" -o bin/aws-assume-role-linux cmd/aws-assume-role/*.go

build-osx:
	@ GOOS=darwin CGO_ENABLED=0 go build -ldflags="$(LDFLAGS)" -o bin/aws-assume-role-osx cmd/aws-assume-role/*.go

build-osx-arm:
	@ GOOS=darwin GOARCH=arm64 CGO_ENABLED=0 go build -ldflags="$(LDFLAGS)" -o bin/aws-assume-role-osx-arm cmd/aws-assume-role/*.go
