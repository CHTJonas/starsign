SHELL := bash
.ONESHELL:

VER=$(shell git describe --tags --always --dirty)
GO=$(shell which go)
GOMOD=$(GO) mod
GOFMT=$(GO) fmt
GOBUILD=$(GO) build -mod=readonly -ldflags "-X main.version=$(VER)"

dir:
	@if [ ! -d bin ]; then mkdir -p bin; fi

mod:
	@$(GOMOD) download

format:
	@$(GOFMT) ./...

build/linux/armv7: dir mod
	export CGO_ENABLED=0
	export GOOS=linux
	export GOARCH=arm
	export GOARM=7
	$(GOBUILD) -o bin/starsign-linux-$(VER:v%=%)-armv7 cmd/starsign/*.go

build/linux/arm64: dir mod
	export CGO_ENABLED=0
	export GOOS=linux
	export GOARCH=arm64
	$(GOBUILD) -o bin/starsign-linux-$(VER:v%=%)-arm64 cmd/starsign/*.go

build/linux/386: dir mod
	export CGO_ENABLED=0
	export GOOS=linux
	export GOARCH=386
	$(GOBUILD) -o bin/starsign-linux-$(VER:v%=%)-386 cmd/starsign/*.go

build/linux/amd64: dir mod
	export CGO_ENABLED=0
	export GOOS=linux
	export GOARCH=amd64
	$(GOBUILD) -o bin/starsign-linux-$(VER:v%=%)-amd64 cmd/starsign/*.go

build/linux: build/linux/armv7 build/linux/arm64 build/linux/386 build/linux/amd64

build/freebsd/armv7: dir mod
	export CGO_ENABLED=0
	export GOOS=freebsd
	export GOARCH=arm
	export GOARM=7
	$(GOBUILD) -o bin/starsign-freebsd-$(VER:v%=%)-armv7 cmd/starsign/*.go

build/freebsd/arm64: dir mod
	export CGO_ENABLED=0
	export GOOS=freebsd
	export GOARCH=arm64
	$(GOBUILD) -o bin/starsign-freebsd-$(VER:v%=%)-arm64 cmd/starsign/*.go

build/freebsd/386: dir mod
	export CGO_ENABLED=0
	export GOOS=freebsd
	export GOARCH=386
	$(GOBUILD) -o bin/starsign-freebsd-$(VER:v%=%)-386 cmd/starsign/*.go

build/freebsd/amd64: dir mod
	export CGO_ENABLED=0
	export GOOS=freebsd
	export GOARCH=amd64
	$(GOBUILD) -o bin/starsign-freebsd-$(VER:v%=%)-amd64 cmd/starsign/*.go

build/freebsd: build/freebsd/armv7 build/freebsd/arm64 build/freebsd/386 build/freebsd/amd64

build/openbsd/armv7: dir mod
	export CGO_ENABLED=0
	export GOOS=openbsd
	export GOARCH=arm
	export GOARM=7
	$(GOBUILD) -o bin/starsign-openbsd-$(VER:v%=%)-armv7 cmd/starsign/*.go

build/openbsd/arm64: dir mod
	export CGO_ENABLED=0
	export GOOS=openbsd
	export GOARCH=arm64
	$(GOBUILD) -o bin/starsign-openbsd-$(VER:v%=%)-arm64 cmd/starsign/*.go

build/openbsd/386: dir mod
	export CGO_ENABLED=0
	export GOOS=openbsd
	export GOARCH=386
	$(GOBUILD) -o bin/starsign-openbsd-$(VER:v%=%)-386 cmd/starsign/*.go

build/openbsd/amd64: dir mod
	export CGO_ENABLED=0
	export GOOS=openbsd
	export GOARCH=amd64
	$(GOBUILD) -o bin/starsign-openbsd-$(VER:v%=%)-amd64 cmd/starsign/*.go

build/openbsd: build/openbsd/armv7 build/openbsd/arm64 build/openbsd/386 build/openbsd/amd64

build/darwin/amd64: dir mod
	export CGO_ENABLED=0
	export GOOS=darwin
	export GOARCH=amd64
	$(GOBUILD) -o bin/starsign-darwin-$(VER:v%=%)-amd64 cmd/starsign/*.go

build/darwin: build/darwin/amd64

build/windows/386: dir mod
	export CGO_ENABLED=0
	export GOOS=windows
	export GOARCH=386
	$(GOBUILD) -o bin/starsign-windows-$(VER:v%=%)-386 cmd/starsign/*.go

build/windows/amd64: dir mod
	export CGO_ENABLED=0
	export GOOS=windows
	export GOARCH=amd64
	$(GOBUILD) -o bin/starsign-windows-$(VER:v%=%)-amd64 cmd/starsign/*.go

build/windows: build/windows/386 build/windows/amd64

build: build/linux build/freebsd build/openbsd build/darwin build/windows

clean:
	@rm -rf bin

all: format build
