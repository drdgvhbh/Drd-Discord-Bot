package repository

import "drdgvhbh/discordbot/internal/entity"

type StockRepository interface {
	UpsertStock(stock *entity.Stock) (*entity.Stock, error)
}
