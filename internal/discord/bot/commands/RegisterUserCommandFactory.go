package commands

import (
	"drdgvhbh/discordbot/internal/user/domain"
	"drdgvhbh/discordbot/internal/user/entity"
	"log"

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
			err := userRepository.InsertUser(user)

			if err != nil {
				if _, ok := err.(*domain.DuplicateUserInsertion); ok {
					write("You are already registered!")
				} else {
					log.Fatalln(err)
				}
			} else {
				write("You are now registered.")
			}

			return nil
		},
	}
}
