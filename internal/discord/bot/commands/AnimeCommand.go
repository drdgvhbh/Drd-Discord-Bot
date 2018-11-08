package commands

import "github.com/urfave/cli"

func CreateAnimeCommand(
	subCommands []cli.Command,
) *cli.Command {
	return &cli.Command{
		Name:        "anime",
		Usage:       "List all anime related commands",
		Subcommands: subCommands,
	}
}

var animeCommand *cli.Command

func ProvideAnimeCommand(
	subCommands []cli.Command,
) *cli.Command {
	if animeCommand != nil {
		return animeCommand
	}

	animeCommand = CreateAnimeCommand(subCommands)

	return animeCommand
}
