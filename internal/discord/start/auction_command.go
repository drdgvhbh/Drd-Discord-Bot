package start

import (
	"drdgvhbh/discordbot/internal/model"
	"drdgvhbh/discordbot/internal/start"

	"github.com/urfave/cli"
)

func NewAuctionCommand(
	startAuctionService start.AuctionService,
) *cli.Command {
	return &cli.Command{
		Name:  "new",
		Usage: "Start a new auction",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "server, s",
				Usage: "discord server identifier",
			},
		},
		Action: func(ctx *cli.Context) error {
			serverID := model.NewDiscordServer(ctx.String("server"))

			animeCharacter := model.NewAnimeCharacter("derp", "ruin")
			startAuctionService.StartAuctionFor(animeCharacter, serverID)

			return nil
		},
	}
}
