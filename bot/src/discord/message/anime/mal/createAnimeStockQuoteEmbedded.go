package mal

import (
	"Drd-Discord-Bot/bot/src/cli/anime/mal"
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

type AnimeStockQuoteEmbeddedOptions struct {
	AnimeStock mal.AnimeStock
}

func CreateAnimeStockQuoteEmbedded(
	options AnimeStockQuoteEmbeddedOptions,
) *discordgo.MessageEmbed {
	animeStock := options.AnimeStock
	price := animeStock.MarketPrice()
	imgURL := animeStock.ImageURL()

	message := &discordgo.MessageEmbed{
		Title:  "Anime Stock Quote",
		Author: &discordgo.MessageEmbedAuthor{},
		Color:  0xFFFFCC,
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:   "Price",
				Value:  fmt.Sprintf("%.4f", price),
				Inline: true,
			},
		},
		Image: &discordgo.MessageEmbedImage{
			URL: imgURL,
		},
		Timestamp: time.Now().Format(time.RFC3339),
	}

	return message
}
