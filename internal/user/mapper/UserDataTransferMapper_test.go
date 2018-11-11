package mapper_test

/*

import (
	"drdgvhbh/discordbot/internal/user/api"
	"drdgvhbh/discordbot/internal/user/entity"
	"drdgvhbh/discordbot/internal/user/mapper"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/google/uuid"
)

func TestCreateDTOFrom(t *testing.T) {
	userDataTransferMapper := mapper.CreateUserMapper()

	id := uuid.New().String()
	tokens := 1337.2
	user := entity.CreateUser(id, tokens)

	assert.Equal(
		t,
		userDataTransferMapper.CreateDTOFrom(user),
		&api.User{
			ID:     id,
			Tokens: tokens,
		})
}

func TestCreateUserFrom(t *testing.T) {
	userDataTransferMapper := mapper.CreateUserMapper()

	id := uuid.New().String()
	tokens := 1337.2
	user := &api.User{
		ID:     id,
		Tokens: tokens,
	}

	assert.Equal(
		t,
		userDataTransferMapper.CreateUserFrom(user),
		entity.CreateUser(id, tokens))
}
*/
