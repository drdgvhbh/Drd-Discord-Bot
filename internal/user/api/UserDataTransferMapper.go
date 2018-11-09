package api

import "drdgvhbh/discordbot/internal/user/entity"

type UserDataTransferMapper interface {
	CreateUserFrom(*User) *entity.User
	CreateDTOFrom(*entity.User) *User
}
