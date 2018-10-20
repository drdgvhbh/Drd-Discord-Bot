# Ruin and Grief Discord Bot

## Prerequisites

- [Node.js](https://nodejs.org/en/)
- [Go](https://golang.org/dl/)
- [Make](http://gnuwin32.sourceforge.net/packages/make.htm)
- [Discord](https://discordapp.com/)

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
