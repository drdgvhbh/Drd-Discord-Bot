package api

import "Drd-Discord-Bot/bot/src/user/entity"

type UserMapper interface {
	MapTo(*User) *entity.User
	MapFrom(*entity.User) *User
}
