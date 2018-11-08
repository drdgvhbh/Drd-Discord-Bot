package api_test

import (
	"drdgvhbh/discordbot/internal/user/api"
	"drdgvhbh/discordbot/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstructorShouldNotBeASingleton(t *testing.T) {
	assert := assert.New(t)

	instance := api.CreateUserRepository(&mocks.Connector{}, &mocks.UserMapper{})
	nextInstance := api.CreateUserRepository(&mocks.Connector{}, &mocks.UserMapper{})

	assert.False(instance == nextInstance, "Both instance pointers should not be equal")
}

func TestProvidesShouldBeASingleton(t *testing.T) {
	assert := assert.New(t)

	instance := api.ProvideUserRepository(&mocks.Connector{}, &mocks.UserMapper{})
	nextInstance := api.ProvideUserRepository(&mocks.Connector{}, &mocks.UserMapper{})

	assert.True(instance == nextInstance, "Both instance pointers should be equal")
}
