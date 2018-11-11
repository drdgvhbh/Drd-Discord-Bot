package commands

import (
	"drdgvhbh/discordbot/internal/service/stock"
	"fmt"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/urfave/cli"
)

type AnimeStockQuoteCommandFactory struct {
	stockInteractor stock.StockInteractor
}

func CreateAnimeStockQuoteCommandFactory(
	stockInteractor stock.StockInteractor,
) *AnimeStockQuoteCommandFactory {
	return &AnimeStockQuoteCommandFactory{
		stockInteractor,
	}
}

func ProvideAnimeStockQuoteCommandFactory(
	stockInteractor stock.StockInteractor,
) *AnimeStockQuoteCommandFactory {
	return CreateAnimeStockQuoteCommandFactory(
		stockInteractor,
	)
}

func (
	animeStockQuoteCommandFactory AnimeStockQuoteCommandFactory,
) Construct(
	write func(message string),
	writeEmbed func(message *discordgo.MessageEmbed),
) *cli.Command {
	stockInteractor := animeStockQuoteCommandFactory.stockInteractor

	return &cli.Command{
		Name:  "quote",
		Usage: "Gets the price of the stock",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "anime, a",
				Usage: "Anime",
			},
		},
		Action: func(ctx *cli.Context) error {
			animeTitle := ctx.String("anime")

			if animeTitle == "" {
				cli.ShowCommandHelp(ctx, "quote")

				return nil
			}

			stock, err := stockInteractor.StockQuotePrice(animeTitle)
			if err != nil {
				log.Fatalln(err)
			}

			writeEmbed(&discordgo.MessageEmbed{
				Title:       fmt.Sprintf("%s", stock.ID()),
				Description: "Stock Quote",
				Color:       0xFFFFCC,
				Fields: []*discordgo.MessageEmbedField{
					&discordgo.MessageEmbedField{
						Name:   "Price",
						Value:  fmt.Sprintf("%.4f", stock.Price()),
						Inline: true,
					},
				},
				Thumbnail: &discordgo.MessageEmbedThumbnail{
					URL: stock.ImageURL(),
				},
				Timestamp: stock.LastUpdated().Format(time.RFC3339),
			})

			return nil
		},
	}
}
