package di

import (
	"drdgvhbh/discordbot/internal/db/pg"
	"drdgvhbh/discordbot/internal/service/repository"
	"drdgvhbh/discordbot/internal/service/stock/implementation"
	"drdgvhbh/discordbot/internal/start"
	stockApi "drdgvhbh/discordbot/internal/stock/api"
	"drdgvhbh/discordbot/internal/storage/postgres"
	userApi "drdgvhbh/discordbot/internal/user/api"
	"os"

	"github.com/jackc/pgx"
	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func ProvideLogger() *logrus.Logger {
	if logger != nil {
		return logger
	}

	logger = logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		DisableLevelTruncation: true,
	})
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.DebugLevel)

	return logger
}

var userRepositoryLoggerSingleton *userApi.UserRepositoryLogger

func ProvideUserRepositoryLogger(
	userRepository *userApi.UserRepository,
	logger *logrus.Logger,
) *userApi.UserRepositoryLogger {
	if userRepositoryLoggerSingleton != nil {
		return userRepositoryLoggerSingleton
	}

	userRepositoryLoggerSingleton = userApi.CreateUserRepositoryLogger(
		userRepository,
		logger)
	return userRepositoryLoggerSingleton
}

var userRepositorySingleton *userApi.UserRepository

func ProvideUserRepository(
	databaseConnector pg.Connector,
	userDataTransferMapper userApi.UserDataTransferMapper,
) *userApi.UserRepository {
	if userRepositorySingleton != nil {
		return userRepositorySingleton
	}

	userRepositorySingleton = userApi.CreateUserRepository(
		databaseConnector,
		userDataTransferMapper)
	return userRepositorySingleton
}

var stockRepositoryLoggerSingleton *stockApi.StockRepositoryLogger

func ProvideStockRepositoryLogger(
	stockRepository *stockApi.StockRepository,
	logger *logrus.Logger,
) *stockApi.StockRepositoryLogger {
	if stockRepositoryLoggerSingleton != nil {
		return stockRepositoryLoggerSingleton
	}

	stockRepositoryLoggerSingleton = stockApi.CreateStockRepositoryLogger(
		stockRepository,
		logger)
	return stockRepositoryLoggerSingleton
}

var stockRepositorySingleton *stockApi.StockRepository

func ProvideStockRepository(
	databaseConnector pg.Connector,
	stockDataTransferMapper stockApi.StockDataTransferMapper,
) *stockApi.StockRepository {
	if stockRepositorySingleton != nil {
		return stockRepositorySingleton
	}

	stockRepositorySingleton = stockApi.CreateStockRepository(
		databaseConnector,
		stockDataTransferMapper)
	return stockRepositorySingleton
}

var stockInteractorSingleton *implementation.StockInteractorImp

func ProvideStockInteractor(
	stockRepository repository.StockRepository,
	animeRepository repository.AnimeRepository,
) *implementation.StockInteractorImp {
	return implementation.CreateStockInteractorImp(
		&implementation.CreateStockInteractorImpParams{
			StockRepository: stockRepository,
			AnimeRepository: animeRepository,
		})
}

var outputChannel chan interface{}

func ProvideOutputChannel() chan interface{} {
	if outputChannel != nil {
		return outputChannel
	}

	outputChannel = make(chan interface{})

	return outputChannel
}

var connection *pgx.Conn

func ProvideDBConnection() *pgx.Conn {
	if connection != nil {
		return connection
	}

	connection = postgres.NewConnection()
	return connection
}

var repository_ *postgres.Repository

func ProvideRepository(dbConnection *pgx.Conn) *postgres.Repository {
	if repository_ != nil {
		return repository_
	}

	repository_ = postgres.NewRepository(dbConnection)

	return repository_
}

var startAuctionService start.AuctionService

func ProvideStartAuctionService(
	eventChannel chan interface{},
	repository *postgres.Repository,
) start.AuctionService {
	if startAuctionService != nil {
		return startAuctionService
	}

	startAuctionService = start.NewAuctionService(eventChannel, repository)

	return startAuctionService
}
