package commands

import "github.com/urfave/cli"

func NewCurrentAuctionCommand(
	subCommands []cli.Command,
) *cli.Command {
	return &cli.Command{
		Name:        "current",
		Usage:       "Shows the current auction",
		Subcommands: subCommands,
	}
}
