GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
OUTPUT_DIR=../../bin/
ENTRY=./bot.go

BINARY_NAME=$(OUTPUT_DIR)DiscordBot.exe
BINARY_UNIX=$(BINARY_NAME)_unix

build:
	$(GOBUILD) -o $(BINARY_NAME) -v $(ENTRY)

run:
	make build
	$(BINARY_NAME)

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

