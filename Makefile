BUILD_CMD=make build

BIN_DIR=bin
CLI_PKG_DIR=src/cli
DISCORD_PKG_DIR=src/discord

build_discord:
	cd $(DISCORD_PKG_DIR)
	$(BUILD_CMD)

build:
	cp .env bin/
	make build_discord

run:
	cd $(BIN_DIR)
	RuinAndGrief.exe

clean:
	rm -r -f $(BIN_DIR)
