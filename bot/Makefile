GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOGEN=$(GOCMD) generate

BUILD_CMD=make build

BIN_DIR=bin
SRC=./src
BINARY_NAME=$(BIN_DIR)/DiscordBot

build:
	mkdir -p bin
	wire $(SRC)/di
	$(GOBUILD) -o $(BINARY_NAME) -v $(SRC)/main.go

run:
	cd $(BIN_DIR)
	DiscordBot

clean:
	$(GOCLEAN)
	rm -r -f $(BIN_DIR)
	