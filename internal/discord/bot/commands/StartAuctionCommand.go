package commands

import "github.com/urfave/cli"

func NewStartAuctionCommand(
	subCommands []cli.Command,
) *cli.Command {
	return &cli.Command{
		Name:        "new",
		Usage:       "Start a new auction",
		Subcommands: subCommands,
	}
}
