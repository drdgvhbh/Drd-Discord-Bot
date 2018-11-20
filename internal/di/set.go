package di

import (
	animeApi "drdgvhbh/discordbot/internal/anime/anime/api"
	animeMapper "drdgvhbh/discordbot/internal/anime/anime/mapper"
	characterApi "drdgvhbh/discordbot/internal/anime/character/api"
	characterMapper "drdgvhbh/discordbot/internal/anime/character/mapper"
	"drdgvhbh/discordbot/internal/cli"
	"drdgvhbh/discordbot/internal/db/pg"
	"drdgvhbh/discordbot/internal/discord/bot"
	"drdgvhbh/discordbot/internal/discord/bot/commands"
	"drdgvhbh/discordbot/internal/service/repository"
	"drdgvhbh/discordbot/internal/service/stock"
	"drdgvhbh/discordbot/internal/service/stock/implementation"
	stockApi "drdgvhbh/discordbot/internal/stock/api"
	stockMapper "drdgvhbh/discordbot/internal/stock/mapper"
	userApi "drdgvhbh/discordbot/internal/user/api"
	userDomain "drdgvhbh/discordbot/internal/user/domain"
	userMapper "drdgvhbh/discordbot/internal/user/mapper"

	"github.com/google/go-cloud/wire"
)

var pgSet = wire.NewSet(pg.ProvideConfig, pg.ProvideConnector)

var cliSet = wire.NewSet(cli.ProvideConfig, cli.ProvideCLI)

var botSet = wire.NewSet(bot.ProvideConfig, bot.ProvideDiscordBot)

var loggerSet = wire.NewSet(ProvideLogger)

var interactorSet = wire.NewSet(
	ProvideStockInteractor,
	wire.Bind(new(stock.StockInteractor), new(implementation.StockInteractorImp)),
)

var commandSet = wire.NewSet(
	commands.CreateRegisterUserCommandFactory,
	commands.ProvideAnimeStockQuoteCommandFactory,
	commands.ProvideAnimeCommand,
)

var userSet = wire.NewSet(
	ProvideUserRepository,
	ProvideUserRepositoryLogger,
	userMapper.CreateUserMapper,
	wire.Bind(new(userDomain.UserRepository), new(userApi.UserRepositoryLogger)),
	wire.Bind(new(userApi.UserDataTransferMapper), new(userMapper.UserDataTransferMapper)))

var stockSet = wire.NewSet(
	ProvideStockRepository,
	ProvideStockRepositoryLogger,
	stockMapper.CreateStockMapper,
	wire.Bind(new(repository.StockRepository), new(stockApi.StockRepositoryLogger)),
	wire.Bind(new(stockApi.StockDataTransferMapper), new(stockMapper.StockDataTransferMapper)))

var animeSet = wire.NewSet(
	animeApi.ProvideAnimeRepository,
	animeMapper.ProvideAnimeMapper,
	wire.Bind(new(repository.AnimeRepository), new(animeApi.AnimeRepository)),
	wire.Bind(new(animeApi.AnimeMapper), new(animeMapper.AnimeMapper)),
)

var characterSet = wire.NewSet(
	characterApi.ProvideCharacterRepository,
	characterMapper.ProvideCharacterMapper,
	wire.Bind(new(repository.CharacterRepository), new(characterApi.CharacterRepository)),
	wire.Bind(new(characterApi.CharacterMapper), new(characterMapper.CharacterMapper)),
)

var RootSet = wire.NewSet(
	pgSet, cliSet, botSet, loggerSet,
	commandSet, characterSet, animeSet,
	stockSet, userSet,
	interactorSet)

/////
var RootSet2 = wire.NewSet(persistanceSet, domainSet)

var persistanceSet = wire.NewSet(ProvideDBConnection, ProvideRepository)

var domainSet = wire.NewSet(ProvideStartAuctionService, ProvideOutputChannel)
