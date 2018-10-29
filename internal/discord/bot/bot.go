package bot

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/davecgh/go-spew/spew"
	"github.com/pkg/errors"
)

type DiscordBot = discordgo.Session

func CreateDiscordBot(config *Config) *DiscordBot {
	bot, err := discordgo.New(fmt.Sprintf("Bot %s", config.BotToken))

	if err != nil {
		err = errors.Wrapf(err,
			"Failed to create discord session (%s)",
			spew.Sdump(config))
		panic(err)
	}

	return bot
}

var discordBot *DiscordBot

func ProvideDiscordBot(config *Config) *DiscordBot {
	if discordBot != nil {
		return discordBot
	}

	discordBot = CreateDiscordBot(config)

	return discordBot
}
