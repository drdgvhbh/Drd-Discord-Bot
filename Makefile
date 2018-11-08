GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOGEN=$(GOCMD) generate
GORUN=$(GOCMD) run

BIN=./bin
SCRIPTS=./scripts
BINARY_NAME=drdbot
OUTPUT_NAME=$(BIN)/$(BINARY_NAME)

.PHONY: config build run clean

config:
	$(SCRIPTS)/init.sh

build:
	$(GORUN) ./vendor/github.com/google/go-cloud/wire/cmd/wire ./internal/di
	$(GOBUILD) -o $(OUTPUT_NAME) -v cmd/main.go

docker-build:
	docker-compose build

docker-run:
	docker-compose up

test:
	go test ./...

cover: 
	go test ./.../... -coverprofile=cover.out
	go tool cover -html=cover.out -o cover.html

run:
	$(OUTPUT_NAME)

watch:
	realize start

clean:
	rm -r -f $(BIN)/*
