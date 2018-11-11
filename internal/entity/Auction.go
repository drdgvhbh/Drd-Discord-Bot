package entity

import (
	"errors"
	"math"
	"time"

	"github.com/google/uuid"
)

type Auction struct {
	id     uuid.UUID
	endsAt time.Time
	bids   []*Bid
	item   *Stock
}

func NewAuction(
	item *Stock,
) *Auction {
	return &Auction{
		id:     uuid.New(),
		endsAt: time.Now().Add(time.Hour * 7),
		bids:   []*Bid{},
		item:   item,
	}
}

func (auction *Auction) ID() uuid.UUID {
	return auction.id
}

func (auction *Auction) EndsAt() time.Time {
	return auction.endsAt
}

func (auction *Auction) Bids() []*Bid {
	return auction.bids
}

func (auction *Auction) Item() *Stock {
	return auction.item
}

func (auction *Auction) HasEnded() bool {
	return auction.endsAt.Before(time.Now())
}

func (auction *Auction) winningBid() *Bid {
	var winningBid *Bid
	for _, bid := range auction.bids {
		if winningBid == nil {
			winningBid = bid
		} else if winningBid.Value() < bid.Value() {
			winningBid = bid
		}
	}

	return winningBid
}

func (auction *Auction) minimumBidValue() float64 {
	return math.Max(
		auction.item.Price(),
		auction.winningBid().Value())
}

func (auction *Auction) Place(bid *Bid) error {
	if auction.HasEnded() {
		return errors.New("auction has ended")
	}

	if bid.Value() <= auction.minimumBidValue() {
		return errors.New("bid value is too low")
	}

	bid.timeOfBid = time.Now()
	auction.bids = append(auction.bids, bid)

	return nil
}
