package domain

import (
	"drdgvhbh/discordbot/internal/entity"
)

type UserRepository interface {
	InsertUser(user *entity.User) error
}
