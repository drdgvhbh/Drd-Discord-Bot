GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOGEN=$(GOCMD) generate

BIN=./bin
SCRIPTS=./scripts
BINARY_NAME=drdbot
OUTPUT_NAME=$(BIN)/$(BINARY_NAME)

.PHONY: config build run clean

config:
	$(SCRIPTS)/init.sh

build:
	$(GOGEN) ./internal/di
	$(GOBUILD) -o $(OUTPUT_NAME) -v cmd/main.go

run:
	$(OUTPUT_NAME)

clean:
	rm -r -f $(BIN)/*
