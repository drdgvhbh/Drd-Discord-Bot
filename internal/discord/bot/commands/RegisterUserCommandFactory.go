package commands

import (
	"drdgvhbh/discordbot/internal/user/domain"
	"drdgvhbh/discordbot/internal/user/entity"

	"github.com/urfave/cli"
)

type RegisterUserCommandFactory struct {
	userRepository domain.UserRepository
	user           *entity.User
}

func CreateRegisterUserCommandFactory(
	userRepository domain.UserRepository,
	user *entity.User,
) *RegisterUserCommandFactory {
	return &RegisterUserCommandFactory{
		userRepository: userRepository,
		user:           user,
	}
}

func (registerCommandFactory RegisterUserCommandFactory) Construct() *cli.Command {
	return &cli.Command{
		Name:  "register",
		Usage: "Registers a user for this bot's services on a Discord service",
		Action: func(ctx *cli.Context) error {
			userRepository := registerCommandFactory.userRepository
			user := registerCommandFactory.user

			userRepository.InsertUser(user)
			return nil
		},
	}
}
