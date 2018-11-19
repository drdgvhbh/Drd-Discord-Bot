package api

import (
	"drdgvhbh/discordbot/internal/db/pg"
	"drdgvhbh/discordbot/internal/entity"
)

type AuctionRepository struct {
	databaseConnector pg.Connector
}

func NewAuctionRepository(
	databaseConnector pg.Connector,
) *AuctionRepository {
	return &AuctionRepository{
		databaseConnector,
	}
}

func (auctionRepository *AuctionRepository) UpsertAuction(
	auction *entity.Auction,
) (*entity.Auction, error) {
	// databaseConnector := auctionRepository.databaseConnector

	return nil, nil
}
