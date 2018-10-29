package di

import (
	"Drd-Discord-Bot/bot/src/cli"
	"Drd-Discord-Bot/bot/src/db/pg"
	"Drd-Discord-Bot/bot/src/user/api"
	"Drd-Discord-Bot/bot/src/user/domain"
	"Drd-Discord-Bot/bot/src/user/mapper"

	"github.com/google/go-cloud/wire"
)

var pgSet = wire.NewSet(pg.ProvideConfig, pg.ProvideConnector)

var cliSet = wire.NewSet(cli.ProvideConfig)

var userSet = wire.NewSet(
	api.CreateUserRepository,
	mapper.CreateUserMapper,
	wire.Bind(new(domain.UserRepository), new(api.UserRepository)),
	wire.Bind(new(api.UserMapper), new(mapper.UserMapper)))

var RootSet = wire.NewSet(pgSet, cliSet, userSet)
