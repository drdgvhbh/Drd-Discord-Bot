package stock

import "drdgvhbh/discordbot/internal/entity"

type StockInteractor interface {
	StockQuotePrice(animeName string) (*entity.Stock, error)
}
