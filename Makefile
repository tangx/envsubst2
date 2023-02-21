
GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)

APPNAME ?= envsubst2

build: tidy
	go build -o out/$(APPNAME)-$(GOOS)-$(GOARCH) .

buildx:
	GOOS=linux GOARCH=amd64 make build
	GOOS=linux GOARCH=arm64 make build
	GOOS=darwin GOARCH=amd64 make build
	GOOS=darwin GOARCH=arm64 make build

install:
	go install

tidy:
	go mod tidy

clean:
	rm -rf out/

