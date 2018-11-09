package api_test

import (
	"drdgvhbh/discordbot/internal/user/api"
	"drdgvhbh/discordbot/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstructorShouldNotBeASingleton(t *testing.T) {
	assert := assert.New(t)

	instance := api.CreateUserRepository(&mocks.Connector{}, &mocks.UserDataTransferMapper{})
	nextInstance := api.CreateUserRepository(&mocks.Connector{}, &mocks.UserDataTransferMapper{})

	assert.False(instance == nextInstance, "Both instance pointers should not be equal")
}

func TestProvidesShouldBeASingleton(t *testing.T) {
	assert := assert.New(t)

	instance := api.ProvideUserRepository(&mocks.Connector{}, &mocks.UserDataTransferMapper{})
	nextInstance := api.ProvideUserRepository(&mocks.Connector{}, &mocks.UserDataTransferMapper{})

	assert.True(instance == nextInstance, "Both instance pointers should be equal")
}
