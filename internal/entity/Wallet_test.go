package entity_test

import (
	"drdgvhbh/discordbot/internal/entity"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ANewWallet struct {
	suite.Suite
	assert      *assert.Assertions
	basicWallet *entity.Wallet
}

func (suite *ANewWallet) SetupTest() {
	suite.assert = assert.New(suite.T())

	ownerID := uuid.New().String()
	startingBalance := 1000.0
	suite.basicWallet = entity.NewWallet(ownerID, startingBalance)
}

func (suite *ANewWallet) TestHasAllItsFundsAvaliable() {
	ownerID := uuid.New().String()
	startingBalance := 1000.0
	wallet := entity.NewWallet(ownerID, startingBalance)

	suite.assert.Equal(startingBalance, wallet.AvaliableFunds())
}

func (suite *ANewWallet) TestProperlyAllocatesFunds() {
	wallet := suite.basicWallet
	avaliableFunds := wallet.AvaliableFunds()
	fundsToBeAllocated := avaliableFunds * 0.10

	wallet.AllocateFunds(fundsToBeAllocated)
	suite.assert.Equal(
		avaliableFunds-fundsToBeAllocated,
		wallet.AvaliableFunds())
}
func TestANewWallet(t *testing.T) {
	suite.Run(t, new(ANewWallet))
}

type OverAllocatedWallet struct {
	suite.Suite
	assert             *assert.Assertions
	startingBalance    float64
	fundsToBeAllocated float64
	allocationError    error
	wallet             *entity.Wallet
}

func (suite *OverAllocatedWallet) SetupTest() {
	suite.assert = assert.New(suite.T())

	ownerID := uuid.New().String()
	suite.startingBalance = 100.0
	suite.fundsToBeAllocated = suite.startingBalance * 1.01
	suite.wallet = entity.NewWallet(ownerID, suite.startingBalance)

	suite.allocationError = suite.wallet.AllocateFunds(
		suite.fundsToBeAllocated)
}

func (suite *OverAllocatedWallet) TestLeavesAvaliableBalanceUnchanged() {
	suite.assert.Equal(
		suite.startingBalance,
		suite.wallet.AvaliableFunds())
}

func (suite *OverAllocatedWallet) TestReturnsNotEnoughFundsError() {
	suite.assert.NotNil(suite.allocationError)
	suite.assert.IsType(
		new(entity.NotEnoughFunds),
		suite.allocationError)
}

func TestOverAllocatedWallet(t *testing.T) {
	suite.Run(t, new(OverAllocatedWallet))
}
