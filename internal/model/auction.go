package model

import (
	"time"

	"github.com/google/uuid"
)

type Auction struct {
	id       *uuid.UUID
	endsAt   time.Time
	bids     []*Bid
	item     *AnimeCharacter
	location *DiscordServer
}

type NewAuctionParams struct {
	ID       *uuid.UUID
	EndsAt   *time.Time
	Bids     *[]*Bid
	Item     *AnimeCharacter
	Location *DiscordServer
}

func NewAuction(params *NewAuctionParams) *Auction {
	newUUID := uuid.New()
	id := &newUUID
	if params.ID != nil {
		id = params.ID
	}

	endsAt := time.Now().Add(time.Hour * 7)
	if params.EndsAt != nil {
		endsAt = *params.EndsAt
	}

	bids := []*Bid{}
	if params.Bids != nil {
		bids = *params.Bids
	}

	return &Auction{
		id:       id,
		endsAt:   endsAt,
		bids:     bids,
		item:     params.Item,
		location: params.Location,
	}
}

func (auction *Auction) Item() *AnimeCharacter {
	return auction.item
}

func (auction *Auction) ID() uuid.UUID {
	return *auction.id
}

func (auction *Auction) Location() *DiscordServer {
	return auction.location
}

func (auction *Auction) EndsAt() time.Time {
	return auction.endsAt
}

func (auction *Auction) Bids() []*Bid {
	return auction.bids
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

/* func (auction *Auction) minimumBidValue() float64 {
	return math.Max(
		auction.item.Price(),
		auction.winningBid().Value())
}
*/
/* func (auction *Auction) Place(bid *Bid) error {
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
*/
