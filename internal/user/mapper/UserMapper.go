package mapper

import (
	"drdgvhbh/discordbot/internal/user/api"
	"drdgvhbh/discordbot/internal/user/entity"
)

type UserMapper struct {
}

var userMapperInstance *UserMapper

func CreateUserMapper() *UserMapper {
	if userMapperInstance != nil {
		return userMapperInstance
	}
	userMapperInstance = &UserMapper{}

	return userMapperInstance
}

func (userMapper UserMapper) MapTo(user *api.User) *entity.User {
	return entity.CreateUser(user.ID, user.Tokens)
}

func (userMapper UserMapper) MapFrom(user *entity.User) *api.User {
	return &api.User{
		ID:     user.ID(),
		Tokens: user.Tokens(),
	}
}
