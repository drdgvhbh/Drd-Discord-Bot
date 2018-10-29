// +build wireinject

package di

import (
	"drdgvhbh/discordbot/internal/cli"
	"drdgvhbh/discordbot/internal/user/api"

	"github.com/google/go-cloud/wire"
)

func InitializeUserRepository() (*api.UserRepository, error) {
	wire.Build(RootSet)
	return &api.UserRepository{}, nil
}

func InitializeCLI() *cli.CLIApp {
	panic(wire.Build(RootSet))
}
