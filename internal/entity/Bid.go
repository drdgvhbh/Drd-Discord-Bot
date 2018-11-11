package entity

import (
	"time"
)

type Bid struct {
	bidderWallet *Wallet
	timeOfBid    time.Time
	value        float64
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
	bid := &Bid{
		bidderWallet: wallet,
		value:        0,
	}

	err := bid.RaiseBy(initialBid)

	if err != nil {
		return nil, err
	}

	return bid, nil
}
