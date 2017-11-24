.PHONY: build test generate clean

EXAMPLE_NAME ?= aor-entity-example
EXAMPLE_SPECS ?= ./tests/swagger.json
EXAMPLE_DIR ?= $(CURDIR)/generated/$(EXAMPLE_NAME)

BINARY_DIR   ?= $(CURDIR)/bin
BINARY_NAME  ?= gin-swagger-aor
SOURCES      = $(shell find . -name '*.go')
GOPKGS       = $(shell go list ./... | grep -v /vendor/)
TEMPLATES    = $(shell find templates/ -type f -name '*.gotmpl')
BUILD_FLAGS  ?= -v
LDFLAGS      ?= -X main.version=$(VERSION) -w -s

default: build

deps: glide
	go get -v -u github.com/go-swagger/go-swagger/cmd/swagger
	go get -v github.com/Masterminds/glide
	go get -v github.com/gin-contrib/pprof

glide:
	glide install --strip-vendor

clean:
	@rm $(BINARY_DIR)/$(BINARY_NAME)

test:
	go test -v $(GOPKGS)

generate: bindata.go

bindata.go: config.yaml $(TEMPLATES)
	go generate .

install:
	go install 

build: $(BINARY_NAME)

$(BINARY): bindata.go $(SOURCES)
	CGO_ENABLED=0 go build -o $(BINARY_DIR)/$(BINARY_NAME) $(BUILD_FLAGS) -ldflags "$(LDFLAGS)"

example:
	rm -fR $(EXAMPLE_DIR)
	./$(BINARY_NAME) --application $(EXAMPLE_NAME) -f $(EXAMPLE_SPECS) --output-dir $(EXAMPLE_DIR)
	go build -o $(BINARY_DIR)/$(EXAMPLE_NAME) $(EXAMPLE_DIR)/cmd/$(EXAMPLE_NAME)-server/