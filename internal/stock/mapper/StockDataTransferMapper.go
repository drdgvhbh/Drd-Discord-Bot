package mapper

import (
	"drdgvhbh/discordbot/internal/entity"
	"drdgvhbh/discordbot/internal/stock/api"
	"time"
)

type StockDataTransferMapper struct {
}

func CreateStockMapper() *StockDataTransferMapper {
	return &StockDataTransferMapper{}
}

func (
	stockDataTransferMapper StockDataTransferMapper,
) CreateStockFrom(stock *api.Stock) *entity.Stock {
	return entity.NewStock(
		&entity.NewStockParams{
			stock.ID,
			stock.Price,
			stock.LastUpdated,
			stock.ImageURL,
		})
}

func (
	stockDataTransferMapper StockDataTransferMapper,
) CreateDTOFrom(stock *entity.Stock) *api.Stock {
	now := time.Now()
	rounded := time.Date(
		now.Year(),
		now.Month(),
		now.Day(),
		0, 0, 0, 0,
		now.Location())

	return &api.Stock{
		ID:          stock.ID(),
		Price:       stock.Price(),
		LastUpdated: rounded,
		ImageURL:    stock.ImageURL(),
	}
}
