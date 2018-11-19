package service

import (
	"drdgvhbh/discordbot/internal/entity"
)

type AuctionService interface {
	startAuction(*entity.Auction)
}
