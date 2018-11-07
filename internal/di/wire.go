// +build wireinject

package di

import (
	characterApi "drdgvhbh/discordbot/internal/anime/character/api"
	"drdgvhbh/discordbot/internal/cli"
	"drdgvhbh/discordbot/internal/discord/bot"
	"drdgvhbh/discordbot/internal/discord/bot/commands"
	userApi "drdgvhbh/discordbot/internal/user/api"

	"github.com/google/go-cloud/wire"
)

func InitializeUserRepository() *userApi.UserRepository {
	wire.Build(RootSet)
	return &userApi.UserRepository{}
}

func InitializeCharacterRepository() *characterApi.CharacterRepository {
	wire.Build(RootSet)
	return &characterApi.CharacterRepository{}
}

func InitializeCLI() *cli.CLIApp {
	panic(wire.Build(RootSet))
}

func InitializeDiscordBot() *bot.DiscordBot {
	panic(wire.Build(RootSet))
}

func InitializeRegisterUserCommandFactory() *commands.RegisterUserCommandFactory {
	panic(wire.Build(RootSet))
}
