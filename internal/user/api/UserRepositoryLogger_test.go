package api_test

/*

import (
	"drdgvhbh/discordbot/internal/user/api"
	"drdgvhbh/discordbot/internal/user/entity"
	"drdgvhbh/discordbot/mocks"
	"errors"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type userRepositoryLoggerConstructor struct {
	suite.Suite
}

func (constructor *userRepositoryLoggerConstructor) TestShouldNotBeASingleton() {
	assert := assert.New(constructor.T())
	logger, _ := test.NewNullLogger()

	userRepositoryLogger := api.CreateUserRepositoryLogger(
		&mocks.UserRepository{}, logger)
	anotherUserRepositoryLogger := api.CreateUserRepositoryLogger(
		&mocks.UserRepository{}, logger)

	assert.True(userRepositoryLogger != anotherUserRepositoryLogger)
}

func TestUserRepositoryLoggerConstructorSuite(t *testing.T) {
	suite.Run(t, new(userRepositoryLoggerConstructor))
}

type userRepositoryLoggerInsertion struct {
	suite.Suite
	logger               *log.Logger
	loggerHook           *test.Hook
	userRepository       *mocks.UserRepository
	userRepositoryLogger *api.UserRepositoryLogger
}

func (insertion *userRepositoryLoggerInsertion) SetupTest() {
	insertion.logger, insertion.loggerHook = test.NewNullLogger()
	insertion.userRepository = &mocks.UserRepository{}

	logger := insertion.logger
	logger.SetFormatter(&log.JSONFormatter{})
	logger.SetLevel(log.DebugLevel)
	userRepository := insertion.userRepository

	insertion.userRepositoryLogger = api.CreateUserRepositoryLogger(
		userRepository,
		logger)
}

func (
	testSuite *userRepositoryLoggerInsertion,
) TestShouldDelegateToDecoratedUserRepository() {
	userRepository := testSuite.userRepository
	userRepositoryLogger := testSuite.userRepositoryLogger

	user := &entity.User{}
	userRepository.Mock.On("InsertUser", user).Return(nil)

	userRepositoryLogger.InsertUser(user)

	userRepository.AssertCalled(testSuite.T(), "InsertUser", user)
}

func (
	testSuite *userRepositoryLoggerInsertion,
) TestShouldReturnTheDecoratedError() {
	assert := assert.New(testSuite.T())
	userRepository := testSuite.userRepository
	userRepositoryLogger := testSuite.userRepositoryLogger

	decoratedError := errors.New("error!")
	user := &entity.User{}
	userRepository.Mock.On("InsertUser", user).Return(decoratedError)

	returnedError := userRepositoryLogger.InsertUser(user)

	assert.Equal(decoratedError, returnedError)
}

func (
	testSuite *userRepositoryLoggerInsertion,
) TestShouldLogTheParameters() {
	assert := assert.New(testSuite.T())

	logger := testSuite.logger
	loggerHook := testSuite.loggerHook

	userRepository := &mocks.UserRepository{}
	userRepositoryLogger := api.CreateUserRepositoryLogger(
		userRepository,
		logger)
	user := &entity.User{}

	userRepository.Mock.On("InsertUser", user).Return(nil)

	userRepositoryLogger.InsertUser(user)

	assert.Condition(func() bool { return len(loggerHook.Entries) >= 1 })
	assert.Equal(log.DebugLevel, loggerHook.LastEntry().Level)
	assert.Equal(
		"Inserting new user into repository",
		loggerHook.LastEntry().Message)
	assert.Equal(log.Fields{
		"id":     user.ID(),
		"tokens": user.Tokens(),
	}, loggerHook.LastEntry().Data)
}
func (
	testSuite *userRepositoryLoggerInsertion,
) TestShouldLogInsertionErrors() {
	assert := assert.New(testSuite.T())

	logger := testSuite.logger
	loggerHook := testSuite.loggerHook

	userRepository := &mocks.UserRepository{}
	userRepositoryLogger := api.CreateUserRepositoryLogger(
		userRepository,
		logger)
	user := &entity.User{}

	insertionError := errors.New("Insertion Error")
	userRepository.Mock.On("InsertUser", user).Return(insertionError)

	userRepositoryLogger.InsertUser(user)

	assert.Condition(func() bool { return len(loggerHook.Entries) >= 2 })
	assert.Equal(log.ErrorLevel, loggerHook.LastEntry().Level)
	assert.Equal(
		"Failed to insert new user into repository",
		loggerHook.LastEntry().Message)
}

func TestUserRepositoryLoggerInsertionSuite(t *testing.T) {
	suite.Run(t, new(userRepositoryLoggerInsertion))
}
*/
