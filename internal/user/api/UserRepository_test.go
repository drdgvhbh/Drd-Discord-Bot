package api_test

/*

import (
	"drdgvhbh/discordbot/internal/user/api"
	"drdgvhbh/discordbot/internal/user/domain"
	"drdgvhbh/discordbot/internal/user/entity"
	"drdgvhbh/discordbot/mocks"
	"testing"

	pq "github.com/lib/pq"
	"github.com/stretchr/testify/mock"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type userRepositoryInsertion struct {
	suite.Suite
	userDataTransferMapper *mocks.UserDataTransferMapper
	databaseConnector      *mocks.Connector
	insertedRow            *mocks.Row
	userRepository         *api.UserRepository
}

func (insertion *userRepositoryInsertion) SetupTest() {
	insertion.userDataTransferMapper = &mocks.UserDataTransferMapper{}
	insertion.userDataTransferMapper.Mock.On(
		"CreateDTOFrom",
		mock.Anything,
	).Return(&api.User{})

	insertion.insertedRow = &mocks.Row{}
	insertion.databaseConnector = &mocks.Connector{}
	insertion.databaseConnector.Mock.On(
		"QueryRow",
		mock.Anything,
	).Return(insertion.insertedRow)

	insertion.userRepository = api.CreateUserRepository(
		insertion.databaseConnector,
		insertion.userDataTransferMapper)
}

func (
	insertion *userRepositoryInsertion,
) TestShouldReturnDuplicateUserInsertionWhenThereIsAUniqueViolation() {
	assert := assert.New(insertion.T())
	insertionError := &pq.Error{
		Code: "23505",
	}

	insertedRow := insertion.insertedRow
	insertedRow.Mock.On("Scan").Return(insertionError)

	userRepository := insertion.userRepository

	user := &entity.User{}
	err := userRepository.InsertUser(user)
	if assert.Error(err) {
		assert.Equal(
			domain.CreateDuplicateUserInsertionError(user),
			err)
	}
}

func (
	insertion *userRepositoryInsertion,
) TestShouldNotHaveAnErrorUponSuccessfulInsertion() {
	assert := assert.New(insertion.T())

	insertedRow := insertion.insertedRow
	insertedRow.Mock.On("Scan").Return(nil)

	userRepository := insertion.userRepository

	user := &entity.User{}
	err := userRepository.InsertUser(user)

	assert.Nil(err)
}

func (
	insertion *userRepositoryInsertion,
) TestShouldReturnOriginalPGErrorIfDoesNotHandleIt() {
	assert := assert.New(insertion.T())

	const INTERNAL_ERROR_CODE = "XX000"
	insertionError := &pq.Error{
		Code: INTERNAL_ERROR_CODE,
	}

	insertedRow := insertion.insertedRow
	insertedRow.Mock.On("Scan").Return(insertionError)

	userRepository := insertion.userRepository

	user := &entity.User{}
	err := userRepository.InsertUser(user)
	if assert.Error(err) {
		assert.Equal(insertionError, err)
	}
}

func TestUserRepositoryInsertSuite(t *testing.T) {
	suite.Run(t, new(userRepositoryInsertion))
}
*/
