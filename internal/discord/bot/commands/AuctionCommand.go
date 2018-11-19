package commands

import "github.com/urfave/cli"

func NewAuctionCommand(
	subCommands []cli.Command,
) *cli.Command {
	return &cli.Command{
		Name:        "auction",
		Usage:       "List all auction related commands",
		Subcommands: subCommands,
	}
}
