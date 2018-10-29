package domain

import "Drd-Discord-Bot/bot/src/user/entity"

type UserRepository interface {
	InsertUser(user *entity.User)
}
