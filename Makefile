GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOGEN=$(GOCMD) generate

BIN=./bin
SCRIPTS=./scripts
BINARY_NAME=drdbot

.PHONY: config build run clean

config:
	$(SCRIPTS)/init.sh

build:
	$(GOGEN) ./internal/di
	$(GOBUILD) -o ${BIN}/$(BINARY_NAME) -v cmd/main.go

run:
	$(BIN)/$(BINARY_NAME)

clean:
	$(SCRIPTS)/clean.sh $(BINARY_NAME)
