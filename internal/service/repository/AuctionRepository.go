package repository

import (
	"drdgvhbh/discordbot/internal/entity"
)

type AuctionRepository interface {
	UpsertAuction(*entity.Auction) (*entity.Auction, error)
}
