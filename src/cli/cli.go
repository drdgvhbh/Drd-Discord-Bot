package botcli

import _ "github.com/joho/godotenv/autoload"

import (
	"github.com/urfave/cli"
)

func CreateCLI(
	appName string,
	helpName string,
) *cli.App {
	App := cli.NewApp()
	App.Name = appName
	App.Usage = "A general purpose bot for ruining and griefing"
	App.HelpName = helpName

	return App
}
