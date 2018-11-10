package api

import (
	"drdgvhbh/discordbot/internal/user/domain"
	"drdgvhbh/discordbot/internal/user/entity"

	log "github.com/sirupsen/logrus"
)

type UserRepositoryLogger struct {
	decoratedUserRepository domain.UserRepository
	logger                  *log.Logger
}

func CreateUserRepositoryLogger(
	userRepository domain.UserRepository,
	logger *log.Logger,
) *UserRepositoryLogger {
	return &UserRepositoryLogger{userRepository, logger}
}

func (
	userRepositoryLogger UserRepositoryLogger,
) InsertUser(user *entity.User) error {
	decoratedUserRepository := userRepositoryLogger.decoratedUserRepository
	logger := userRepositoryLogger.logger

	logger.WithFields(log.Fields{
		"id":     user.ID(),
		"tokens": user.Tokens(),
	}).Debug("Inserting new user into repository")

	insertionError := decoratedUserRepository.InsertUser(user)

	if insertionError != nil {
		logger.WithError(insertionError).Error(
			"Failed to insert new user into repository")
	}

	return insertionError
}
