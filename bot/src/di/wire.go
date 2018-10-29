// +build wireinject

package di

import (
	"Drd-Discord-Bot/bot/src/user/api"

	"github.com/google/go-cloud/wire"
)

func InitializeUserRepository() (*api.UserRepository, error) {
	wire.Build(RootSet)
	return &api.UserRepository{}, nil
}
