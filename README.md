# Ruin and Grief Discord Bot

## Prerequisites

- [Node.js](https://nodejs.org/en/)
- [Go](https://golang.org/dl/)
- [Make](http://gnuwin32.sourceforge.net/packages/make.htm)
- [Discord](https://discordapp.com/)

## Usage

```
NAME:
  ${APP_NAME} - A general purpose bot for ruining and griefing

 USAGE:
  ${CMD_PREFIX} [global options] command [command options] [arguments...]

 VERSION:
  0.0.0

 COMMANDS:
    anime    List anime commands
    help, h  Shows a list of commands or help for one command

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

```
APP_NAME=Ruin and Grief Bot
CMD_PREFIX=!ruin
BOT_TOKEN=XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
EOF_DELIM=EOF
```

### Install Node packages

```sh
npm install
```

## Running

### Terminal 1

```sh
npm run watch
```

### Terminal 2

```sh
./bin/RuinAndGrief.exe
```
