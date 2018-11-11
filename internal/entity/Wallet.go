package entity

type Wallet struct {
	ownerID        string
	funds          float64
	allocatedFunds float64
}

func (wallet *Wallet) Owner() string {
	return wallet.ownerID
}

func (wallet *Wallet) AllocateFunds(amount float64) error {
	if amount > wallet.AvaliableFunds() {
		return NewNotEnoughFundsError(wallet.AvaliableFunds(), amount)
	}

	wallet.allocatedFunds += amount

	return nil
}

func (wallet *Wallet) AvaliableFunds() float64 {
	return wallet.funds - wallet.allocatedFunds
}

func NewWallet(
	ownerID string,
	startingBalance float64,
) *Wallet {
	return &Wallet{
		ownerID:        ownerID,
		funds:          startingBalance,
		allocatedFunds: 0,
	}
}
