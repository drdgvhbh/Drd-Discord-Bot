# Drd Discord Bot

## Prerequisites

- [Go](https://golang.org/dl/)
- [Make](http://gnuwin32.sourceforge.net/packages/make.htm)
- [Discord](https://discordapp.com/)

## Usage

```
NAME:
  Drd Bot - General purpose discord bot

 USAGE:
  !cmd [global options] command [command options] [arguments...]

 VERSION:
  0.0.0

 COMMANDS:
    register  Registers a user for this bot's services on a Discord service
    anime     List all anime related commands
    auction   List all auction related commands
    help, h   Shows a list of commands or help for one command

 GLOBAL OPTIONS:
  --help, -h     show help
  --version, -v  print the version
```

## Setup

1. Create a developer account on the [Discord Developer Portal](https://discordapp.com/developers/applications/)
2. Create an application and a bot
3. Add your bot to your server using https://discordapp.com/oauth2/authorize?client_id={CLIENTID}&scope=bot
4. Save your **bot token** somewhere

## Building

### Environment File

1. Create a `.env` file in the root of your directory and fill it in with the following fields

```sh
CMD_PREFIX=
BOT_TOKEN=
EOF_DELIM=

DB_PORT=
DB_HOST=
DB_USERNAME=
DB_PASSWORD=
DB_NAME=

```

### Terminal 2

```sh
./bin/drdbot
```
