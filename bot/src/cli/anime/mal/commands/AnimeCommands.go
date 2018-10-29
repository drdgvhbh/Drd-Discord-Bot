package commands

import (
	"Drd-Discord-Bot/bot/src/cli/anime/mal"
	"Drd-Discord-Bot/bot/src/cli/anime/mal/api"
	messageMal "Drd-Discord-Bot/bot/src/discord/message/anime/mal"
	"log"

	"github.com/bwmarrin/discordgo"
	realCli "github.com/urfave/cli"
)

type CommandCallback = func(output *discordgo.MessageEmbed) (*discordgo.Message, error)

func AddAnimeCommands(
	cli *realCli.App,
	callback CommandCallback,
) {
	cli.Commands = append(cli.Commands, realCli.Command{
		Name:  "anime",
		Usage: "List anime commands",
		Action: func(c *realCli.Context) error {
			realCli.ShowCommandHelp(c, "anime")

			return nil
		},
		Subcommands: realCli.Commands{
			getAnimeProfileCommand(callback),
			listAnimeCharacterStock(callback),
		},
	})
}

var getAnimeProfileCommand = func(callback CommandCallback) realCli.Command {
	return realCli.Command{
		Name:  "profile",
		Usage: "Displays a user's anime profile",
		Flags: []realCli.Flag{
			realCli.StringFlag{
				Name:  "name, n",
				Usage: "My Anime List Username",
			},
		},
		Action: func(c *realCli.Context) error {
			userProfile, err := api.GetProfile(c.String("name"))
			if err != nil {
				log.Println(err)
				return nil
			}
			options := messageMal.CreateAnimeProfileEmbeddedOptions{
				AnimeProfile: userProfile,
			}
			embeddedMessage := messageMal.CreateAnimeProfileEmbedded(
				options)

			_, error := callback(embeddedMessage)
			if error != nil {
				log.Println(err)
			}
			return nil
		},
	}
}

var listAnimeCharacterStock = func(callback CommandCallback) realCli.Command {
	return realCli.Command{
		Name:  "quote",
		Usage: "Gets a stock price listing for an anime + anime character pair",
		Flags: []realCli.Flag{
			realCli.IntFlag{
				Name:  "anime, a",
				Usage: "Anime ID",
			},
			realCli.IntFlag{
				Name:  "character, c",
				Usage: "Character ID",
			},
		},
		Action: func(c *realCli.Context) error {
			animeID := c.Int("anime")
			characterID := c.Int("character")

			animeStock, err := mal.CreateAnimeStock(characterID, animeID)

			if err != nil {
				log.Println(err)
				return nil
			}

			options := messageMal.AnimeStockQuoteEmbeddedOptions{
				AnimeStock: animeStock,
			}
			embeddedMessage := messageMal.CreateAnimeStockQuoteEmbedded(options)

			_, error := callback(embeddedMessage)
			if error != nil {
				log.Println(err)
			}
			return nil
		},
	}
}
