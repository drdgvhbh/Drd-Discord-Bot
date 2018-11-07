package commands

import (
	animeDomain "drdgvhbh/discordbot/internal/anime/anime/domain"
	characterDomain "drdgvhbh/discordbot/internal/anime/character/domain"
	"fmt"
	"math"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/urfave/cli"
)

type AnimeStockQuoteCommandFactory struct {
	animeCharacterRepository characterDomain.CharacterRepository
	animeRepository          animeDomain.AnimeRepository
}

func CreateAnimeStockQuoteCommandFactory(
	animeCharacterRepository characterDomain.CharacterRepository,
	animeRepository animeDomain.AnimeRepository,
) *AnimeStockQuoteCommandFactory {
	return &AnimeStockQuoteCommandFactory{
		animeCharacterRepository: animeCharacterRepository,
		animeRepository:          animeRepository,
	}
}

func ProvideAnimeStockQuoteCommandFactory(
	animeCharacterRepository characterDomain.CharacterRepository,
	animeRepository animeDomain.AnimeRepository,
) *AnimeStockQuoteCommandFactory {
	return CreateAnimeStockQuoteCommandFactory(
		animeCharacterRepository, animeRepository)
}

func (
	animeStockQuoteCommandFactory AnimeStockQuoteCommandFactory,
) Construct(
	write func(message string),
	writeEmbed func(message *discordgo.MessageEmbed),
) *cli.Command {
	animeCharacterRepository := animeStockQuoteCommandFactory.animeCharacterRepository
	animeRepository := animeStockQuoteCommandFactory.animeRepository

	return &cli.Command{
		Name:  "quote",
		Usage: "Gets the price of the stock",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "character, c",
				Usage: "Anime Character",
			},
			cli.StringFlag{
				Name:  "anime, a",
				Usage: "Anime",
			},
		},
		Action: func(ctx *cli.Context) error {
			animeCharacterName := ctx.String("character")
			animeTitle := ctx.String("anime")

			if animeCharacterName == "" || animeTitle == "" {
				cli.ShowCommandHelp(ctx, "quote")

				return nil
			}

			characters, characterError := animeCharacterRepository.SearchCharactersByName(
				animeCharacterName)

			animes, animeError := animeRepository.SearchAnimesByTitle(
				animeTitle)

			if characterError != nil || animeError != nil {
				write(fmt.Sprintf("Failed to get quote for %s", animeCharacterName))
				return characterError
			}

			marketPrice := func(score float64) float64 {
				return (math.Pow(score, 3.0) + math.Pow(score, 2.0)) + score
			}

			writeEmbed(&discordgo.MessageEmbed{
				Title:       fmt.Sprintf("%s (%s)", characters[0].Name(), animes[0].Title()),
				Description: "Stock Quote",
				Color:       0xFFFFCC,
				Fields: []*discordgo.MessageEmbedField{
					&discordgo.MessageEmbedField{
						Name:   "Price",
						Value:  fmt.Sprintf("%.4f", marketPrice(animes[0].Score())),
						Inline: true,
					},
				},
				Timestamp: time.Now().Format(time.RFC3339),
			})

			return nil
		},
	}
}
