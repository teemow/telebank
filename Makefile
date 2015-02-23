PROJECT=telebank
ORGANIZATION=teemow

SOURCE := $(shell find . -name '*.go')
VERSION := $(shell cat VERSION)
GOPATH := $(shell pwd)/.gobuild
PROJECT_PATH := $(GOPATH)/src/github.com/$(ORGANIZATION)

.PHONY=all clean test deps bin

all: deps $(PROJECT)

clean:
	rm -rf $(GOPATH) $(PROJECT) simple

test:
	GOPATH=$(GOPATH) go test ./...

# deps
deps: .gobuild
.gobuild:
	mkdir -p $(PROJECT_PATH)
	cd $(PROJECT_PATH) && ln -s ../../../.. $(PROJECT)

	#
	# Fetch public packages
	GOPATH=$(GOPATH) go get -d github.com/$(ORGANIZATION)/$(PROJECT)

	#
	# Fetch test packages
	GOPATH=$(GOPATH) go get -d github.com/onsi/gomega
	GOPATH=$(GOPATH) go get -d github.com/onsi/ginkgo

# build
$(PROJECT): $(SOURCE) VERSION
	GOPATH=$(GOPATH) go build -ldflags "-X main.projectVersion $(VERSION)" -o $(PROJECT)

install: $(PROJECT)
	cp $(PROJECT) /usr/local/bin/
