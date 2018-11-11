package api_test

/* package api_test

import (
	"drdgvhbh/discordbot/internal/stock/api"
	"drdgvhbh/discordbot/internal/stock/entity"
	"drdgvhbh/discordbot/mocks"
	"errors"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type stockRepositoryLoggerConstructor struct {
	suite.Suite
}

func (constructor *stockRepositoryLoggerConstructor) TestShouldNotBeASingleton() {
	assert := assert.New(constructor.T())
	logger, _ := test.NewNullLogger()

	stockRepositoryLogger := api.CreateStockRepositoryLogger(
		&mocks.StockRepository{}, logger)
	anotherStockRepositoryLogger := api.CreateStockRepositoryLogger(
		&mocks.StockRepository{}, logger)

	assert.True(stockRepositoryLogger != anotherStockRepositoryLogger)
}

func TestStockRepositoryLoggerConstructorSuite(t *testing.T) {
	suite.Run(t, new(stockRepositoryLoggerConstructor))
}

type stockRepositoryLoggerInsertion struct {
	suite.Suite
	logger                *log.Logger
	loggerHook            *test.Hook
	stockRepository       *mocks.StockRepository
	stockRepositoryLogger *api.StockRepositoryLogger
}

func (insertion *stockRepositoryLoggerInsertion) SetupTest() {
	insertion.logger, insertion.loggerHook = test.NewNullLogger()
	insertion.stockRepository = &mocks.StockRepository{}

	logger := insertion.logger
	logger.SetFormatter(&log.JSONFormatter{})
	logger.SetLevel(log.DebugLevel)
	stockRepository := insertion.stockRepository

	insertion.stockRepositoryLogger = api.CreateStockRepositoryLogger(
		stockRepository,
		logger)
}

func (
	testSuite *stockRepositoryLoggerInsertion,
) TestShouldDelegateToDecoratedStockRepository() {
	stockRepository := testSuite.stockRepository
	stockRepositoryLogger := testSuite.stockRepositoryLogger

	stock := &entity.Stock{}
	stockRepository.Mock.On("UpsertStock", stock).Return(nil, nil)

	stockRepositoryLogger.UpsertStock(stock)

	stockRepository.AssertCalled(testSuite.T(), "UpsertStock", stock)
}

func (
	testSuite *stockRepositoryLoggerInsertion,
) TestShouldReturnTheDecoratedError() {
	assert := assert.New(testSuite.T())
	stockRepository := testSuite.stockRepository
	stockRepositoryLogger := testSuite.stockRepositoryLogger

	decoratedError := errors.New("error!")
	stock := &entity.Stock{}
	stockRepository.Mock.On("UpsertStock", stock).Return(nil, decoratedError)

	_, returnedError := stockRepositoryLogger.UpsertStock(stock)

	assert.Equal(decoratedError, returnedError)
}

func (
	testSuite *stockRepositoryLoggerInsertion,
) TestShouldLogTheParameters() {
	assert := assert.New(testSuite.T())

	logger := testSuite.logger
	loggerHook := testSuite.loggerHook

	stockRepository := &mocks.StockRepository{}
	stockRepositoryLogger := api.CreateStockRepositoryLogger(
		stockRepository,
		logger)
	stock := &entity.Stock{}

	stockRepository.Mock.On(
		"UpsertStock", stock,
	).Return(
		nil, nil,
	)

	stockRepositoryLogger.UpsertStock(stock)

	assert.Condition(func() bool { return len(loggerHook.Entries) >= 1 })
	assert.Equal(log.DebugLevel, loggerHook.LastEntry().Level)
	assert.Equal(
		"Upserting stock into repository",
		loggerHook.LastEntry().Message)
	assert.Equal(log.Fields{
		"id":    stock.ID(),
		"price": stock.Price(),
	}, loggerHook.LastEntry().Data)
}
func (
	testSuite *stockRepositoryLoggerInsertion,
) TestShouldLogInsertionErrors() {
	assert := assert.New(testSuite.T())

	logger := testSuite.logger
	loggerHook := testSuite.loggerHook

	stockRepository := &mocks.StockRepository{}
	stockRepositoryLogger := api.CreateStockRepositoryLogger(
		stockRepository,
		logger)
	stock := &entity.Stock{}

	insertionError := errors.New("Insertion Error")
	stockRepository.Mock.On(
		"UpsertStock", stock,
	).Return(
		nil, insertionError,
	)

	stockRepositoryLogger.UpsertStock(stock)

	assert.Condition(func() bool { return len(loggerHook.Entries) >= 2 })
	assert.Equal(log.ErrorLevel, loggerHook.LastEntry().Level)
	assert.Equal(
		"Failed to upsert stock into repository",
		loggerHook.LastEntry().Message)
}

func TestStockRepositoryLoggerInsertionSuite(t *testing.T) {
	suite.Run(t, new(stockRepositoryLoggerInsertion))
}
*/
