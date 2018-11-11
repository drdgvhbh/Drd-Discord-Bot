package mapper_test

/*

import (
	"drdgvhbh/discordbot/internal/stock/api"
	"drdgvhbh/discordbot/internal/stock/entity"
	"drdgvhbh/discordbot/internal/stock/mapper"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/google/uuid"
)

func TestCreateDTOFrom(t *testing.T) {
	stockDataTransferMapper := mapper.CreateStockMapper()

	id := uuid.New().String()
	price := 56.2
	stock := entity.NewStock(id, price)
	now := time.Now()
	rounded := time.Date(
		now.Year(),
		now.Month(),
		now.Day(),
		0, 0, 0, 0,
		now.Location())
	assert.Equal(
		t,
		stockDataTransferMapper.CreateDTOFrom(stock),
		&api.Stock{
			ID:          id,
			Price:       price,
			LastUpdated: rounded,
		})
}

func TestCreateStockFrom(t *testing.T) {
	stockDataTransferMapper := mapper.CreateStockMapper()

	id := uuid.New().String()
	price := 75.2
	stock := &api.Stock{
		ID:    id,
		Price: price,
	}

	assert.Equal(
		t,
		stockDataTransferMapper.CreateStockFrom(stock),
		entity.NewStock(id, price))
}
*/
