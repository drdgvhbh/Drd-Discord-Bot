package di

import (
	characterApi "drdgvhbh/discordbot/internal/anime/character/api"
	characterDomain "drdgvhbh/discordbot/internal/anime/character/domain"
	characterMapper "drdgvhbh/discordbot/internal/anime/character/mapper"
	"drdgvhbh/discordbot/internal/cli"
	"drdgvhbh/discordbot/internal/db/pg"
	"drdgvhbh/discordbot/internal/discord/bot"
	"drdgvhbh/discordbot/internal/discord/bot/commands"
	userApi "drdgvhbh/discordbot/internal/user/api"
	userDomain "drdgvhbh/discordbot/internal/user/domain"
	userMapper "drdgvhbh/discordbot/internal/user/mapper"

	"github.com/google/go-cloud/wire"
)

var pgSet = wire.NewSet(pg.ProvideConfig, pg.ProvideConnector)

var cliSet = wire.NewSet(cli.ProvideConfig, cli.ProvideCLI)

var botSet = wire.NewSet(bot.ProvideConfig, bot.ProvideDiscordBot)

var commandSet = wire.NewSet(
	commands.CreateRegisterUserCommandFactory,
	commands.ProvideAnimeStockQuoteCommandFactory,
)

var userSet = wire.NewSet(
	userApi.CreateUserRepository,
	userMapper.CreateUserMapper,
	wire.Bind(new(userDomain.UserRepository), new(userApi.UserRepository)),
	wire.Bind(new(userApi.UserMapper), new(userMapper.UserMapper)))

var characterSet = wire.NewSet(
	characterApi.ProvideCharacterRepository,
	characterMapper.ProvideCharacterMapper,
	wire.Bind(new(characterDomain.CharacterRepository), new(characterApi.CharacterRepository)),
	wire.Bind(new(characterApi.CharacterMapper), new(characterMapper.CharacterMapper)),
)

var RootSet = wire.NewSet(pgSet, cliSet, botSet, commandSet, characterSet, userSet)
