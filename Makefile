VERSION      := $(shell cat VERSION)
BUILD        := $(shell git rev-parse --short HEAD)

LDFLAGS       = -ldflags "-X=main.version=$(VERSION) -X=main.build=$(BUILD)"

GOCMD         = go
GOBUILD       = $(GOCMD) build $(LDFLAGS)
GOCLEAN       = $(GOCMD) clean

BINARY_NAME   = mtproto_proxy_exporter
BINARY_386    = $(BINARY_NAME)_386
BINARY_AMD64  = $(BINARY_NAME)_amd64

all: deps build

build:
	$(GOBUILD) -o $(BINARY_NAME)

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_386)
	rm -f $(BINARY_AMD64)

run:
	$(GOCMD) run $(BINARY_NAME).go

deps:
	dep ensure

build-linux-386:
	CGO_ENABLED=0 GOOS=linux GOARCH=386 $(GOBUILD) -o $(BINARY_386)

build-linux-amd64:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_AMD64)