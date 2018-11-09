package domain

import (
	"drdgvhbh/discordbot/internal/user/entity"
	"fmt"
)

type DuplicateUserInsertion struct {
	user *entity.User
}

func (duplicateUserInsertion DuplicateUserInsertion) Error() string {
	return fmt.Sprintf(
		"user with id  %s has already been inserted",
		duplicateUserInsertion.user.ID())
}
