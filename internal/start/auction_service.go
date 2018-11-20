package start

import (
	"drdgvhbh/discordbot/internal/model"
)

type AnimeCharacter = model.AnimeCharacter
type DiscordServer = model.DiscordServer
type Auction = model.Auction
type NewAuctionParams = model.NewAuctionParams

type AuctionService interface {
	StartAuctionFor(
		animeCharacter *AnimeCharacter,
		server *DiscordServer)
}

type Repository interface {
	AddAuction(auction *Auction) error
}

type auctionService struct {
	eventChannel chan interface{}
	repository   Repository
}

func NewAuctionService(
	eventChannel chan interface{},
	repository Repository,
) AuctionService {
	return &auctionService{
		eventChannel: eventChannel,
		repository:   repository,
	}
}

func (auctionService *auctionService) StartAuctionFor(
	animeCharacter *AnimeCharacter,
	server *DiscordServer,
) {
	newAuction := model.NewAuction(&NewAuctionParams{
		Item:     animeCharacter,
		Location: server,
	})

	repository := auctionService.repository
	eventChannel := auctionService.eventChannel
	err := repository.AddAuction(newAuction)
	if err != nil {
		go func() {
			eventChannel <- AuctionEvent{
				auction:      newAuction,
				isSuccessful: false,
				errors: []error{
					err,
				},
			}
		}()

		return
	}

	go func() {
		eventChannel <- AuctionEvent{
			auction:      newAuction,
			isSuccessful: true,
			errors:       []error{},
		}
	}()
}

type AuctionEvent struct {
	auction      *Auction
	isSuccessful bool
	errors       []error
}

func (auctionEvent *AuctionEvent) Auction() *Auction {
	return auctionEvent.auction
}

func (auctionEvent *AuctionEvent) IsSuccessful() bool {
	return auctionEvent.isSuccessful
}

func (auctionEvent *AuctionEvent) Errors() []error {
	errors := []error{}
	copy(auctionEvent.errors, errors)
	return errors
}
