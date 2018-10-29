package cli

import _ "github.com/joho/godotenv/autoload"

import (
	"github.com/urfave/cli"
)

func CreateCLI(
	config *Config,
) *cli.App {
	App := cli.NewApp()
	App.Name = config.Name()
	App.Usage = config.Description()
	App.HelpName = config.HelpPrefix()

	return App
}
