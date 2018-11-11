package implementation

import (
	"drdgvhbh/discordbot/internal/entity"
	"drdgvhbh/discordbot/internal/service/repository"
	"drdgvhbh/discordbot/internal/service/stock"
	"math"
)

type StockInteractorImp struct {
	stockRepository repository.StockRepository
	animeRepository repository.AnimeRepository
}

type CreateStockInteractorImpParams struct {
	StockRepository repository.StockRepository
	AnimeRepository repository.AnimeRepository
}

func CreateStockInteractorImp(
	params *CreateStockInteractorImpParams,
) *StockInteractorImp {
	return &StockInteractorImp{
		params.StockRepository,
		params.AnimeRepository,
	}
}

func (
	stockInteractor StockInteractorImp,
) StockQuotePrice(
	animeName string,
) (*entity.Stock, error) {
	animes, err := stockInteractor.animeRepository.SearchAnimesByTitle(
		animeName,
	)

	marketPrice := func(score float64) float64 {
		return (math.Pow(score, 3.0) + math.Pow(score, 2.0)) + score
	}

	if len(animes) == 0 {
		return nil, stock.CreateStockDoesNotExistError(animeName)
	}

	firstAnimeResult := animes[0]

	price := marketPrice(firstAnimeResult.Score())

	stockRepository := stockInteractor.stockRepository
	insertedStock, insertionError := stockRepository.UpsertStock(
		entity.NewStock(
			&entity.NewStockParams{
				ID:       firstAnimeResult.Title(),
				Price:    price,
				ImageURL: firstAnimeResult.ImageURL(),
			}))
	if insertionError != nil {
		return nil, insertionError
	}

	return insertedStock, err
}
