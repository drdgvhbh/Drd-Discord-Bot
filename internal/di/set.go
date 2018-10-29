package di

import (
	"drdgvhbh/discordbot/internal/cli"
	"drdgvhbh/discordbot/internal/db/pg"
	"drdgvhbh/discordbot/internal/discord/bot"
	"drdgvhbh/discordbot/internal/user/api"
	"drdgvhbh/discordbot/internal/user/domain"
	"drdgvhbh/discordbot/internal/user/mapper"

	"github.com/google/go-cloud/wire"
)

var pgSet = wire.NewSet(pg.ProvideConfig, pg.ProvideConnector)

var cliSet = wire.NewSet(cli.ProvideConfig, cli.ProvideCLI)

var botSet = wire.NewSet(bot.ProvideConfig, bot.ProvideDiscordBot)

var userSet = wire.NewSet(
	api.CreateUserRepository,
	mapper.CreateUserMapper,
	wire.Bind(new(domain.UserRepository), new(api.UserRepository)),
	wire.Bind(new(api.UserMapper), new(mapper.UserMapper)))

var RootSet = wire.NewSet(pgSet, cliSet, botSet, userSet)
