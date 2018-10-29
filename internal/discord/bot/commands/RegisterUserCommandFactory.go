package commands

import (
	"drdgvhbh/discordbot/internal/user/domain"
	"drdgvhbh/discordbot/internal/user/entity"

	"github.com/urfave/cli"
)

type RegisterUserCommandFactory struct {
	userRepository domain.UserRepository
	user           *entity.User
	write          func(message string)
}

func CreateRegisterUserCommandFactory(
	userRepository domain.UserRepository,
	user *entity.User,
	write func(message string),
) *RegisterUserCommandFactory {
	return &RegisterUserCommandFactory{
		userRepository: userRepository,
		user:           user,
		write:          write,
	}
}

func (registerCommandFactory RegisterUserCommandFactory) Construct() *cli.Command {
	write := registerCommandFactory.write
	userRepository := registerCommandFactory.userRepository
	user := registerCommandFactory.user

	return &cli.Command{
		Name:  "register",
		Usage: "Registers a user for this bot's services on a Discord service",
		Action: func(ctx *cli.Context) error {
			userRepository.InsertUser(user)
			write("You are now registered.")

			return nil
		},
	}
}
