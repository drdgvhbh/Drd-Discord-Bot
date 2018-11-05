package domain

import (
	"drdgvhbh/discordbot/internal/user/entity"
	"fmt"
)

type UserAlreadyExistsError struct {
	user *entity.User
}

func (userAlreadyExistsError UserAlreadyExistsError) Error() string {
	return fmt.Sprintf(
		"user with id  %s already exists",
		userAlreadyExistsError.user.ID())
}
