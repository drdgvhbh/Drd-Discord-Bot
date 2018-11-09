package api_test

import (
	"drdgvhbh/discordbot/internal/user/api"
	"drdgvhbh/discordbot/mocks"
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type constructor struct {
	suite.Suite
}

func (constructor *constructor) TestShouldNotBeASingleton() {
	assert := assert.New(constructor.T())

	userRepository := api.CreateUserRepository(&mocks.Connector{}, &mocks.UserDataTransferMapper{})
	anotherUserRepository := api.CreateUserRepository(&mocks.Connector{}, &mocks.UserDataTransferMapper{})

	assert.True(userRepository != anotherUserRepository)
}

func TestConstructorSuite(t *testing.T) {
	suite.Run(t, new(constructor))
}

type provider struct {
	suite.Suite
}

func (provider *provider) TestShouldBeASingleton() {
	assert := assert.New(provider.T())

	userRepository := api.ProvideUserRepository(&mocks.Connector{}, &mocks.UserDataTransferMapper{})
	anotherUserRepository := api.ProvideUserRepository(&mocks.Connector{}, &mocks.UserDataTransferMapper{})

	assert.True(userRepository == anotherUserRepository)
}

func TestProviderSuite(t *testing.T) {
	suite.Run(t, new(provider))
}

type insertion struct {
	suite.Suite
}

func (
	insertion *insertion,
) TestShouldReturnDuplicateUserInsertionWhenThereIsAUniqueViolation() {
	/* 	insertedRow := */

	databaseConnector := &mocks.Connector{}
	databaseConnector.Mock.On("Insert User", mock.Anything)
}
