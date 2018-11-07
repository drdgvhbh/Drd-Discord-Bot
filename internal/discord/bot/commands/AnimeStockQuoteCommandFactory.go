package commands

import (
	characterDomain "drdgvhbh/discordbot/internal/anime/character/domain"

	"github.com/urfave/cli"
)

type AnimeStockQuoteCommandFactory struct {
	animeCharacterRepository characterDomain.CharacterRepository
}

func CreateAnimeStockQuoteCommandFactory(
	animeCharacterRepository characterDomain.CharacterRepository,
) *AnimeStockQuoteCommandFactory {
	return &AnimeStockQuoteCommandFactory{
		animeCharacterRepository: animeCharacterRepository,
	}
}

func (
	animeStockQuoteCommandFactory AnimeStockQuoteCommandFactory,
) Construct(
	animeCharacterName string,
) *cli.Command {
	animeCharacterRepository := animeStockQuoteCommandFactory.animeCharacterRepository

	return &cli.Command{}
}
