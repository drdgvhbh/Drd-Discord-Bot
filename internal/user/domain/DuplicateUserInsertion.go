package domain

import (
	"drdgvhbh/discordbot/internal/user/entity"
	"fmt"
)

type DuplicateUserInsertion struct {
	user *entity.User
}

func CreateDuplicateUserInsertionError(
	user *entity.User,
) *DuplicateUserInsertion {
	return &DuplicateUserInsertion{
		user,
	}
}

func (duplicateUserInsertion DuplicateUserInsertion) Error() string {
	return fmt.Sprintf(
		"user %s has already been inserted",
		duplicateUserInsertion.user.ID())
}
