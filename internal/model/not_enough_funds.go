package model

import (
	"fmt"
)

type NotEnoughFunds struct {
	avaliableFunds float64
	fundsRequired  float64
}

func NewNotEnoughFundsError(
	avaliableFunds float64,
	fundsRequired float64,
) *NotEnoughFunds {
	return &NotEnoughFunds{
		avaliableFunds: avaliableFunds,
		fundsRequired:  fundsRequired,
	}
}

func (notEnoughFunds *NotEnoughFunds) AvaliableFunds() float64 {
	return notEnoughFunds.avaliableFunds
}

func (notEnoughFunds *NotEnoughFunds) FundsRequired() float64 {
	return notEnoughFunds.fundsRequired
}

func (notEnoughFunds *NotEnoughFunds) Error() string {
	return fmt.Sprintf("Not enough funds to process transaction")
}
