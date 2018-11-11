package api_test

/*

import (
	"drdgvhbh/discordbot/internal/stock/api"
	"drdgvhbh/discordbot/internal/stock/entity"
	"drdgvhbh/discordbot/mocks"
	"testing"

	pq "github.com/lib/pq"
	"github.com/stretchr/testify/mock"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type stockRepositoryInsertion struct {
	suite.Suite
	stockDataTransferMapper *mocks.StockDataTransferMapper
	databaseConnector       *mocks.Connector
	stockRepository         *api.StockRepository
}

func (insertion *stockRepositoryInsertion) SetupTest() {
	insertion.stockDataTransferMapper = &mocks.StockDataTransferMapper{}
	insertion.stockDataTransferMapper.Mock.On(
		"CreateDTOFrom",
		mock.Anything,
	).Return(&api.Stock{})

	insertion.databaseConnector = &mocks.Connector{}

	insertion.stockRepository = api.CreateStockRepository(
		insertion.databaseConnector,
		insertion.stockDataTransferMapper)
}

func (insertion *stockRepositoryInsertion) InsertionCall() *mock.Call {
	return insertion.databaseConnector.Mock.On(
		"QueryRow",
		`
			INSERT INTO stocks(id, price)
			VALUES (
				$1, $2
			)
			ON CONFLICT (id)
			DO
				UPDATE
					SET price = EXCLUDED.price,
						last_updated = EXCLUDED.last_updated
			RETURNING
				id, price, last_updated;
		`,
		mock.Anything,
		mock.Anything,
		mock.Anything,
	)
}

func (
	insertion *stockRepositoryInsertion,
) TestShouldReturnTheInsertedStockUponSuccessfulInsertion() {
	assert := assert.New(insertion.T())

	rowQuery := &mocks.Row{}
	rowQuery.Mock.On(
		"Scan", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	insertion.InsertionCall().Return(rowQuery)

	stockRepository := insertion.stockRepository

	stock := &entity.Stock{}
	insertedStock, err := stockRepository.UpsertStock(stock)

	assert.NotNil(insertedStock)
	assert.Nil(err)
}

func (
	insertion *stockRepositoryInsertion,
) TestShouldReturnOriginalPGErrorIfDoesNotHandleIt() {
	assert := assert.New(insertion.T())

	const INTERNAL_ERROR_CODE = "XX000"
	insertionError := &pq.Error{
		Code: INTERNAL_ERROR_CODE,
	}

	rowQuery := &mocks.Row{}
	rowQuery.Mock.On(
		"Scan", mock.Anything, mock.Anything, mock.Anything,
	).Return(insertionError)
	insertion.InsertionCall().Return(rowQuery)

	stockRepository := insertion.stockRepository

	stock := &entity.Stock{}
	_, err := stockRepository.UpsertStock(stock)
	if assert.Error(err) {
		assert.Equal(insertionError, err)
	}
}

func TestStockRepositoryInsertSuite(t *testing.T) {
	suite.Run(t, new(stockRepositoryInsertion))
}
*/
