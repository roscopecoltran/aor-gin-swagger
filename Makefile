.PHONY: build test generate clean

BINARY       ?= gin-swagger-gorm
SOURCES      = $(shell find . -name '*.go')
GOPKGS       = $(shell go list ./... | grep -v /vendor/)
TEMPLATES    = $(shell find templates/ -type f -name '*.gotmpl')
BUILD_FLAGS  ?= -v
LDFLAGS      ?= -X main.version=$(VERSION) -w -s

default: build

deps:
	go get -v -u github.com/go-swagger/go-swagger/cmd/swagger
	go get -v github.com/Masterminds/glide
	glide install --strip-vendor

clean:
	@rm $(BINARY)

test:
	go test -v $(GOPKGS)

generate: bindata.go

bindata.go: config.yaml $(TEMPLATES)
	go generate .

build: $(BINARY)

$(BINARY): bindata.go $(SOURCES)
	CGO_ENABLED=0 go build -o $(BINARY) $(BUILD_FLAGS) -ldflags "$(LDFLAGS)"
