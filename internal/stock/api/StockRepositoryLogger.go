package api

import (
	"drdgvhbh/discordbot/internal/entity"
	"drdgvhbh/discordbot/internal/service/repository"

	log "github.com/sirupsen/logrus"
)

type StockRepositoryLogger struct {
	decoratedStockRepository repository.StockRepository
	logger                   *log.Logger
}

func CreateStockRepositoryLogger(
	stockRepository repository.StockRepository,
	logger *log.Logger,
) *StockRepositoryLogger {
	return &StockRepositoryLogger{stockRepository, logger}
}

func (
	stockRepositoryLogger StockRepositoryLogger,
) UpsertStock(
	stock *entity.Stock,
) (*entity.Stock, error) {
	decoratedStockRepository := stockRepositoryLogger.decoratedStockRepository
	logger := stockRepositoryLogger.logger

	logger.WithFields(log.Fields{
		"id":       stock.ID(),
		"price":    stock.Price(),
		"imageURL": stock.ImageURL(),
	}).Debug("Upserting stock into repository")

	insertedStock, insertionError := decoratedStockRepository.UpsertStock(
		stock)

	if insertionError != nil {
		logger.WithError(insertionError).Error(
			"Failed to upsert stock into repository")
	}

	return insertedStock, insertionError
}
