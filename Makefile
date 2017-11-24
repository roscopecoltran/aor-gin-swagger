.PHONY: build test generate clean

OS          		?= $(shell go env GOOS)
ARCH        		?= $(shell go env GOARCH)

EXAMPLE_NAME 		?= aor-entity-example
EXAMPLE_SPECS 		?= $(CURDIR)/tests/swagger.yaml
EXAMPLE_DIR 		?= $(CURDIR)/generated/$(EXAMPLE_NAME)

SWAGGER_UI_VERSION 	?= 3.5.0

BINARY_DIR   		?= $(CURDIR)/bin
BINARY_NAME  		?= gin-swagger-aor
ifeq ($(OS),windows)
	BINARY_NAME 	:= $(BINARY_NAME).exe
endif

SOURCES      		= $(shell find . -name '*.go')
GOPKGS       		= $(shell go list ./... | grep -v /vendor/)
TEMPLATES    		= $(shell find templates/ -type f -name '*.gotmpl')
BUILD_FLAGS  		?= -v
LDFLAGS      		?= -X main.version=$(VERSION) -w -s

default: build

# git submodule add --prefix ./addons/swagger-ui https://github.com/swagger-api/swagger-ui master --squash
# git submodule add --prefix ./addons/swagger-editor https://github.com/swagger-api/swagger-editor master --squash

.PHONY: install-go-bindata
install-go-bindata:
	which go-bindata 2>/dev/null  || go get -v github.com/jteeuwen/go-bindata/go-bindata

deps: glide
	go get -v -u github.com/roscopecoltran/go-swagger/cmd/swagger
	go get -v github.com/Masterminds/glide
	go get -v github.com/gin-contrib/pprof
	go get -v github.com/NYTimes/openapi2proto

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

$(BINARY_NAME): bindata.go $(SOURCES)
	CGO_ENABLED=0 go build -o $(BINARY_DIR)/$(BINARY_NAME) $(BUILD_FLAGS) -ldflags "$(LDFLAGS)"

quick-example: example example-bin example-run

example:
	rm -fR $(EXAMPLE_DIR)
	$(BINARY_DIR)/$(BINARY_NAME) --application $(EXAMPLE_NAME) -f $(EXAMPLE_SPECS) --output-dir $(EXAMPLE_DIR)

example-bin:
	cd $(EXAMPLE_DIR)/cmd/$(EXAMPLE_NAME)-server && go build -o $(BINARY_DIR)/$(EXAMPLE_NAME)

example-run:
	./bin/$(EXAMPLE_NAME) --insecure-http

build-assets:
	go-bindata -pkg=assets -prefix=assets/swagger-ui -o=./assets/assets.go ./assets/swagger-ui/...

swagger_ui_resources.go: swagger-ui
	mkdir -p $(dir $@)
	go-bindata -nomemcopy -o $@ -prefix ${PWD}/swagger-ui -pkg swaggerui $</...

$(BINARY_NAME)/swaggerui.go: swagger-ui
	mkdir -p $(dir $@)
	go-bindata -nocompress -o $@ -prefix ${PWD} -pkg $(BINARY_NAME) $</...

swagger-ui:
	curl -sSL https://github.com/swagger-api/swagger-ui/archive/v$(SWAGGER_UI_VERSION).zip > /tmp/swagger-ui.zip
	mkdir -p $@
	unzip /tmp/swagger-ui.zip ./addons/swagger-ui-*/dist/* -d $@
	mv swagger-ui/swagger-ui-*/dist/* ./addons/swagger-ui/
	rm -r ./addons/swagger-ui/swagger-ui-*