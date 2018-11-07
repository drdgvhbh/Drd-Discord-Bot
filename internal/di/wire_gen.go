// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package di

import (
	api2 "drdgvhbh/discordbot/internal/anime/character/api"
	mapper2 "drdgvhbh/discordbot/internal/anime/character/mapper"
	cli2 "drdgvhbh/discordbot/internal/cli"
	"drdgvhbh/discordbot/internal/db/pg"
	"drdgvhbh/discordbot/internal/discord/bot"
	"drdgvhbh/discordbot/internal/discord/bot/commands"
	"drdgvhbh/discordbot/internal/user/api"
	"drdgvhbh/discordbot/internal/user/mapper"
	"github.com/bwmarrin/discordgo"
	"github.com/urfave/cli"
)

// Injectors from wire.go:

func InitializeUserRepository() *api.UserRepository {
	config := pg.ProvideConfig()
	db := pg.ProvideConnector(config)
	userMapper := mapper.CreateUserMapper()
	userRepository := api.CreateUserRepository(db, userMapper)
	return userRepository
}

func InitializeCharacterRepository() *api2.CharacterRepository {
	characterMapper := mapper2.ProvideCharacterMapper()
	characterRepository := api2.ProvideCharacterRepository(characterMapper)
	return characterRepository
}

func InitializeCLI() *cli.App {
	config := cli2.ProvideConfig()
	app := cli2.ProvideCLI(config)
	return app
}

func InitializeDiscordBot() *discordgo.Session {
	config := bot.ProvideConfig()
	session := bot.ProvideDiscordBot(config)
	return session
}

func InitializeRegisterUserCommandFactory() *commands.RegisterUserCommandFactory {
	config := pg.ProvideConfig()
	db := pg.ProvideConnector(config)
	userMapper := mapper.CreateUserMapper()
	userRepository := api.CreateUserRepository(db, userMapper)
	registerUserCommandFactory := commands.CreateRegisterUserCommandFactory(userRepository)
	return registerUserCommandFactory
}
