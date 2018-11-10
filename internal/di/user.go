package di

import (
	"drdgvhbh/discordbot/internal/user/api"

	log "github.com/sirupsen/logrus"
)

var userRepositoryLoggerSingleton *api.UserRepositoryLogger

func ProvideUserRepositoryLogger(
	userRepository *api.UserRepository,
	logger *log.Logger,
) *api.UserRepositoryLogger {
	if userRepositoryLoggerSingleton != nil {
		return userRepositoryLoggerSingleton
	}

	userRepositoryLoggerSingleton = api.CreateUserRepositoryLogger(
		userRepository,
		logger)
	return userRepositoryLoggerSingleton
}
