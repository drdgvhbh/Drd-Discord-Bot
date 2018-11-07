package commands

import (
	"drdgvhbh/discordbot/internal/anime/character/domain"
	"fmt"

	"github.com/urfave/cli"
)

type AnimeStockQuoteCommandFactory struct {
	animeCharacterRepository domain.CharacterRepository
}

func CreateAnimeStockQuoteCommandFactory(
	animeCharacterRepository domain.CharacterRepository,
) *AnimeStockQuoteCommandFactory {
	return &AnimeStockQuoteCommandFactory{
		animeCharacterRepository: animeCharacterRepository,
	}
}

func ProvideAnimeStockQuoteCommandFactory(
	animeCharacterRepository domain.CharacterRepository,
) *AnimeStockQuoteCommandFactory {
	return CreateAnimeStockQuoteCommandFactory(animeCharacterRepository)
}

func (
	animeStockQuoteCommandFactory AnimeStockQuoteCommandFactory,
) Construct(
	write func(message string),
) *cli.Command {
	animeCharacterRepository := animeStockQuoteCommandFactory.animeCharacterRepository

	return &cli.Command{
		Name:  "quote",
		Usage: "Gets the price of the stock",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "character, c",
				Usage: "Anime Character",
			},
		},
		Action: func(ctx *cli.Context) error {
			animeCharacterName := ctx.String("character")

			if animeCharacterName == "" {
				cli.ShowCommandHelp(ctx, "quote")

				return nil
			}

			characters, err := animeCharacterRepository.SearchCharactersByName(
				animeCharacterName)

			if err != nil {
				write(fmt.Sprintf("Failed to get quote for %s", animeCharacterName))
				return err
			}
			write(characters[0].Name())

			return nil
		},
	}
}
