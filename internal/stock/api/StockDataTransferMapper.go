package api

import "drdgvhbh/discordbot/internal/entity"

type StockDataTransferMapper interface {
	CreateStockFrom(*Stock) *entity.Stock
	CreateDTOFrom(*entity.Stock) *Stock
}
