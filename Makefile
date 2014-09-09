PROJECT=taskq
ORGANIZATION=zyndiecate

SOURCE := $(shell find . -name '*.go')
GOPATH := $(shell pwd)/.gobuild
PROJECT_PATH := $(GOPATH)/src/github.com/$(ORGANIZATION)

.PHONY=all clean test

all: $(GOPATH) $(PROJECT)

clean:
	rm -rf $(GOPATH) $(PROJECT)

test:
	GOPATH=$(GOPATH) go test ./...

# deps
$(GOPATH):
	mkdir -p $(PROJECT_PATH)
	cd $(PROJECT_PATH) && ln -s ../../../.. $(PROJECT)

	#
	# Fetch private packages

	#
	# Fetch public packages
	GOPATH=$(GOPATH) go get -d github.com/$(ORGANIZATION)/$(PROJECT)

	#
	# Fetch test packages
	GOPATH=$(GOPATH) go get -d github.com/onsi/gomega
	GOPATH=$(GOPATH) go get -d github.com/onsi/ginkgo

# build
$(PROJECT): $(SOURCE)
	GOPATH=$(GOPATH) go build -o $(PROJECT)
