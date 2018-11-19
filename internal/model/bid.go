package model

import (
	"time"

	"github.com/google/uuid"
)

type Bid struct {
	id           *uuid.UUID
	bidderWallet *Wallet
	timeOfBid    time.Time
	value        float64
}

func (bid *Bid) ID() uuid.UUID {
	return *bid.id
}

func (bid *Bid) Value() float64 {
	return bid.value
}

func (bid *Bid) RaiseBy(amount float64) error {
	err := bid.bidderWallet.AllocateFunds(amount)
	if err != nil {
		return err
	}

	bid.value += amount

	return err
}

func NewBid(
	wallet *Wallet,
	initialBid float64,
) (*Bid, error) {
	id := uuid.New()
	bid := &Bid{
		id:           &id,
		bidderWallet: wallet,
		value:        0,
	}

	err := bid.RaiseBy(initialBid)

	if err != nil {
		return nil, err
	}

	return bid, nil
}
