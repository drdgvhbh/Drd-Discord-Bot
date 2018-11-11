package implementation_test

import (
	entity "drdgvhbh/discordbot/internal/entity"
	"drdgvhbh/discordbot/internal/mocks"
	"drdgvhbh/discordbot/internal/service/stock"
	"drdgvhbh/discordbot/internal/service/stock/implementation"
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type stockQuotePriceSuite struct {
	suite.Suite
	stockRepository *mocks.StockRepository
	animeRepository *mocks.AnimeRepository
}

func (suite *stockQuotePriceSuite) SetupTest() {
	suite.stockRepository = &mocks.StockRepository{}
	suite.animeRepository = &mocks.AnimeRepository{}
}

func (
	suite *stockQuotePriceSuite,
) CreateStockInteractor() *implementation.StockInteractorImp {
	return implementation.CreateStockInteractorImp(
		&implementation.CreateStockInteractorImpParams{
			StockRepository: suite.stockRepository,
			AnimeRepository: suite.animeRepository,
		})
}

func (
	suite *stockQuotePriceSuite,
) TestReturnsMarketPriceUsingAnimeScore() {
	assert := assert.New(suite.T())

	const score float64 = 7.5
	const trueMarketPrice float64 = 485.625
	animeName := "Overlord"

	animeRepository := suite.animeRepository
	animeRepository.Mock.On(
		"SearchAnimesByTitle", animeName,
	).Return(
		[]*entity.Anime{
			entity.NewAnime(
				&entity.NewAnimeParams{
					ID:    "654154-645694",
					Title: animeName,
					Score: score,
				},
			),
			// Should only take the score from the first anime entry
			entity.NewAnime(
				&entity.NewAnimeParams{
					ID:    "12546546-26565",
					Title: "Steins's GateName",
					Score: 9.5,
				}),
		},
		nil)

	stockRepository := suite.stockRepository
	stockRepository.Mock.On("UpsertStock", mock.Anything).Return(
		entity.NewStock(&entity.NewStockParams{
			ID:    animeName,
			Price: trueMarketPrice,
		}), nil)

	stockInteractor := suite.CreateStockInteractor()

	stockQuote, err := stockInteractor.StockQuotePrice(animeName)
	assert.Nil(err)
	animeRepository.AssertCalled(
		suite.T(),
		"SearchAnimesByTitle",
		animeName)
	assert.Equal(
		trueMarketPrice,
		stockQuote.Price())
}

func (
	suite *stockQuotePriceSuite,
) TestReturnsStockDoesNotExistErrorIfNoAnimesFound() {
	assert := assert.New(suite.T())

	name := "diasdhfiesgffisuzfu"
	animeRepository := suite.animeRepository
	animeRepository.Mock.On(
		"SearchAnimesByTitle", name,
	).Return(
		[]*entity.Anime{}, nil)

	stockRepository := suite.stockRepository
	stockRepository.Mock.On("UpsertStock", mock.Anything).Return(nil, nil)

	stockInteractor := suite.CreateStockInteractor()

	stockQuote, err := stockInteractor.StockQuotePrice(name)

	assert.Equal(
		stock.CreateStockDoesNotExistError(name),
		err)
	assert.Nil(stockQuote)
}

func (
	suite *stockQuotePriceSuite,
) TestShouldUpsertStockPriceIntoRepository() {
	animeName := "Overlord"
	const score float64 = 1
	animeRepository := suite.animeRepository
	animeRepository.Mock.On(
		"SearchAnimesByTitle", animeName,
	).Return(
		[]*entity.Anime{
			entity.NewAnime(
				&entity.NewAnimeParams{
					ID:    "654154-645694",
					Title: animeName,
					Score: score,
				}),
		}, nil)

	expectedStock := entity.NewStock(
		&entity.NewStockParams{
			ID:    animeName,
			Price: 3,
		})
	stockRepository := suite.stockRepository
	stockRepository.Mock.On("UpsertStock", expectedStock).Return(nil, nil)

	stockInteractor := suite.CreateStockInteractor()
	stockInteractor.StockQuotePrice(animeName)

	stockRepository.AssertCalled(suite.T(), "UpsertStock", expectedStock)
}

func TestStockInteractorSuite(t *testing.T) {
	suite.Run(t, new(stockQuotePriceSuite))
}
