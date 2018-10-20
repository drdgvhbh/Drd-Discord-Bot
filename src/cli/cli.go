package botcli

import _ "github.com/joho/godotenv/autoload"

import (
	"os"

	"github.com/urfave/cli"
)

func CreateCLI() *cli.App {
	App := cli.NewApp()
	App.Name = os.Getenv("APP_NAME")
	App.Usage = "A general purpose bot for ruining and griefing"
	App.HelpName = os.Getenv("HELP_NAME")

	return App
}
