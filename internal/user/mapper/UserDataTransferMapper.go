package mapper

import (
	"drdgvhbh/discordbot/internal/user/api"
	"drdgvhbh/discordbot/internal/user/entity"
)

type UserDataTransferMapper struct {
}

var userMapperInstance *UserDataTransferMapper

func CreateUserMapper() *UserDataTransferMapper {
	if userMapperInstance != nil {
		return userMapperInstance
	}
	userMapperInstance = &UserDataTransferMapper{}

	return userMapperInstance
}

func (userDataTransferMapper UserDataTransferMapper) CreateUserFrom(user *api.User) *entity.User {
	return entity.CreateUser(user.ID, user.Tokens)
}

func (userDataTransferMapper UserDataTransferMapper) CreateDTOFrom(user *entity.User) *api.User {
	return &api.User{
		ID:     user.ID(),
		Tokens: user.Tokens(),
	}
}
