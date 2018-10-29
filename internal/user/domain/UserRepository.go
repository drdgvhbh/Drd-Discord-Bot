package domain

import "drdgvhbh/discordbot/internal/user/entity"

type UserRepository interface {
	InsertUser(user *entity.User)
}
