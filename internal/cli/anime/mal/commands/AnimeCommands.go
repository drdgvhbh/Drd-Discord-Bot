package commands

import (
	"drdgvhbh/discordbot/internal/cli/anime/mal/api"
	messageMal "drdgvhbh/discordbot/internal/discord/message/anime/mal"
	"log"

	"github.com/bwmarrin/discordgo"
	realCli "github.com/urfave/cli"
)

type CommandCallback = func(output *discordgo.MessageEmbed) (*discordgo.Message, error)

var GetAnimeProfileCommand = func(callback CommandCallback) realCli.Command {
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
