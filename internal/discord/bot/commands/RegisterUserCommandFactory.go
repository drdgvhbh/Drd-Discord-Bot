package commands

import (
	"drdgvhbh/discordbot/internal/user/domain"
	"drdgvhbh/discordbot/internal/user/entity"

	"github.com/urfave/cli"
)

type RegisterUserCommandFactory struct {
	userRepository domain.UserRepository
}

func CreateRegisterUserCommandFactory(
	userRepository domain.UserRepository,
) *RegisterUserCommandFactory {
	return &RegisterUserCommandFactory{
		userRepository: userRepository,
	}
}

func (
	registerCommandFactory RegisterUserCommandFactory,
) Construct(
	user *entity.User,
	write func(message string),
) *cli.Command {
	userRepository := registerCommandFactory.userRepository

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
