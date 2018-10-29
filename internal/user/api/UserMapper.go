package api

import "drdgvhbh/discordbot/internal/user/entity"

type UserMapper interface {
	MapTo(*User) *entity.User
	MapFrom(*entity.User) *User
}
