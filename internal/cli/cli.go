package cli

import _ "github.com/joho/godotenv/autoload"

import (
	"github.com/urfave/cli"
)

type CLIApp = cli.App

func CreateCLI(
	config *Config,
) *CLIApp {
	App := cli.NewApp()
	App.Name = config.Name()
	App.Usage = config.Description()
	App.HelpName = config.HelpPrefix()

	return App
}

func ProvideCLI(config *Config) *CLIApp {
	return CreateCLI(config)
}
