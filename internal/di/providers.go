package di

import (
	"drdgvhbh/discordbot/internal/db/pg"
	userApi "drdgvhbh/discordbot/internal/user/api"
	"os"

	"github.com/sirupsen/logrus"
)

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
