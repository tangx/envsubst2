PKG = $(shell cat go.mod | grep "^module " | sed -e "s/module //g")
VERSION = v$(shell cat .version)
COMMIT_SHA ?= $(shell git describe --always)-devel

GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)
GOBUILD=CGO_ENABLED=0 go build -a -ldflags "-X ${PKG}/version.Version=${VERSION}+sha.${COMMIT_SHA}"

APPNAME ?= envsubst2

build: tidy
	$(GOBUILD) -o out/$(APPNAME)-$(GOOS)-$(GOARCH) .

buildx:
	GOOS=linux GOARCH=amd64 make build
	GOOS=linux GOARCH=arm64 make build
	GOOS=darwin GOARCH=amd64 make build
	GOOS=darwin GOARCH=arm64 make build

install: build
	cp out/$(APPNAME)-$(GOOS)-$(GOARCH) $(GOPATH)/bin/$(APPNAME)
	

tidy:
	go mod tidy

clean:
	rm -rf out/

