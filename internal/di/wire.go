// +build wireinject

package di

import (
	characterApi "drdgvhbh/discordbot/internal/anime/character/api"
	"drdgvhbh/discordbot/internal/cli"
	"drdgvhbh/discordbot/internal/discord/bot"
	"drdgvhbh/discordbot/internal/discord/bot/commands"
	stockApi "drdgvhbh/discordbot/internal/stock/api"
	userApi "drdgvhbh/discordbot/internal/user/api"

	nativeCli "github.com/urfave/cli"

	"github.com/google/go-cloud/wire"
)

func InitializeUserRepository() *userApi.UserRepository {
	wire.Build(RootSet)
	return &userApi.UserRepository{}
}

func InitializeStockRepository() *stockApi.StockRepository {
	wire.Build(RootSet)
	return &stockApi.StockRepository{}
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

func InitializeAnimeStockQuoteCommandFactory() *commands.AnimeStockQuoteCommandFactory {
	panic(wire.Build(RootSet))
}

func InitializeAnimeCommand(
	subCommands []nativeCli.Command,
) *nativeCli.Command {
	return commands.ProvideAnimeCommand(subCommands)
}
