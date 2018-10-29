package mapper

import (
	"Drd-Discord-Bot/bot/src/user/api"
	"Drd-Discord-Bot/bot/src/user/entity"
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
	return entity.CreateUser(user.ID)
}

func (userMapper UserMapper) MapFrom(user *entity.User) *api.User {
	return &api.User{
		ID: user.ID(),
	}
}
