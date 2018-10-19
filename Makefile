# Basic go commands
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
ENTRY_DIR=./src/app/

# Binary names
BINARY_NAME=RuinAndGrief.exe
BINARY_UNIX=$(BINARY_NAME)_unix

build:
        $(GOBUILD) -o $(BINARY_NAME) -v $(ENTRY_DIR)

run:
        $(GOBUILD) -o $(BINARY_NAME) -v $(ENTRY_DIR)
        ./$(BINARY_NAME)

test:
  		$(GOTEST) -v ./...

clean:
        $(GOCLEAN)
        rm -f $(BINARY_NAME)
        rm -f $(BINARY_UNIX)

install:
        $(GOGET) github.com/parnurzeal/gorequest
        $(GOGET) github.com/bwmarrin/discordgo
	$(GOGET) github.com/joho/godotenv
        $(GOGET) github.com/urfave/cli
