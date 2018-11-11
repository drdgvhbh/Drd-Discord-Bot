package api

import (
	"drdgvhbh/discordbot/internal/db/pg"
	"drdgvhbh/discordbot/internal/entity"

	pq "github.com/lib/pq"
)

type StockRepository struct {
	databaseConnector       pg.Connector
	stockDataTransferMapper StockDataTransferMapper
}

func CreateStockRepository(
	databaseConnector pg.Connector,
	stockDataTransferMapper StockDataTransferMapper,
) *StockRepository {
	return &StockRepository{databaseConnector, stockDataTransferMapper}
}

func (
	stockRepository StockRepository,
) UpsertStock(
	stock *entity.Stock,
) (*entity.Stock, error) {
	databaseConnector := stockRepository.databaseConnector
	stockDataTransferMapper := stockRepository.stockDataTransferMapper

	stockDTO := stockDataTransferMapper.CreateDTOFrom(stock)

	rowQuery := databaseConnector.QueryRow(
		`
			INSERT INTO stocks(id, price, image_url) 
			VALUES (
				$1, $2, $3
			)
			ON CONFLICT (id) 
			DO 
				UPDATE
					SET price = EXCLUDED.price,
						image_url = EXCLUDED.image_url,
						last_updated = EXCLUDED.last_updated
			RETURNING
				id, price, last_updated, image_url;
		`,
		stockDTO.ID,
		stockDTO.Price,
		stockDTO.ImageURL,
	)

	insertedStock := new(Stock)

	upsertStockError := rowQuery.Scan(
		&insertedStock.ID,
		&insertedStock.Price,
		&insertedStock.LastUpdated,
		&insertedStock.ImageURL)

	if upsertStockError != nil {
		if upsertStockPGError, isPGError := upsertStockError.(*pq.Error); isPGError {
			switch upsertStockPGError.Code.Name() {
			default:
				return nil, upsertStockPGError
			}
		}
	}

	return stockDataTransferMapper.CreateStockFrom(insertedStock), nil
}
