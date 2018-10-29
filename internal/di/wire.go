// +build wireinject

package di

import (
	"drdgvhbh/discordbot/internal/cli"
	"drdgvhbh/discordbot/internal/discord/bot"
	"drdgvhbh/discordbot/internal/discord/bot/commands"
	"drdgvhbh/discordbot/internal/user/api"
	"drdgvhbh/discordbot/internal/user/entity"

	"github.com/google/go-cloud/wire"
)

func InitializeUserRepository() *api.UserRepository {
	wire.Build(RootSet)
	return &api.UserRepository{}
}

func InitializeCLI() *cli.CLIApp {
	panic(wire.Build(RootSet))
}

func InitializeDiscordBot() *bot.DiscordBot {
	panic(wire.Build(RootSet))
}

func InitializeRegisterUserCommandFactory(
	user *entity.User,
) *commands.RegisterUserCommandFactory {
	return commands.CreateRegisterUserCommandFactory(InitializeUserRepository(), user)
}
