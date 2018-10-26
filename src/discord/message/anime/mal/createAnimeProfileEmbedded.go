package mal

import (
	"cli/anime/mal/api/response"
	"time"

	"github.com/bwmarrin/discordgo"
)

type CreateAnimeProfileEmbeddedOptions struct {
	AnimeProfile *response.UserProfileResponse
}

func CreateAnimeProfileEmbedded(
	options CreateAnimeProfileEmbeddedOptions,
) *discordgo.MessageEmbed {
	animeProfile := options.AnimeProfile

	var favoriteAnimes string
	for _, animes := range animeProfile.Favorites.Anime {
		favoriteAnimes += animes.Name
		favoriteAnimes += "\n"
	}
	birthday := animeProfile.Birthday
	gender := animeProfile.Gender
	imageURL := animeProfile.ImageURL

	embed := &discordgo.MessageEmbed{
		Author:      &discordgo.MessageEmbedAuthor{},
		Color:       0x2E51A2, // Green
		Description: "My Anime List Profile",
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:   "Last Online",
				Value:  animeProfile.LastOnline.Format("01/02/2006 15:04"),
				Inline: true,
			},
			&discordgo.MessageEmbedField{
				Name: "Gender",
				Value: func() string {
					if gender != "" {
						return gender
					}
					return "Unknown"
				}(),
				Inline: true,
			},
			&discordgo.MessageEmbedField{
				Name: "Birthday",
				Value: func() string {
					if !birthday.IsZero() {
						return animeProfile.Birthday.Format("01/02/2006")
					}
					return "Unknown"
				}(),
				Inline: true,
			},
			&discordgo.MessageEmbedField{
				Name:   "Joined",
				Value:  animeProfile.Joined.Format("01/02/2006"),
				Inline: true,
			},
			&discordgo.MessageEmbedField{
				Name: "Favorite Animes",
				Value: func() string {
					if favoriteAnimes != "" {
						return favoriteAnimes
					}
					return "N/A"
				}(),
				Inline: false,
			},
		},
		Image: &discordgo.MessageEmbedImage{
			URL: func() string {
				if imageURL != "" {
					return imageURL
				}
				return "https://goo.gl/RqHQVi"

			}(),
		},
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: func() string {
				if imageURL != "" {
					return imageURL
				}
				return "https://goo.gl/RqHQVi"

			}(),
		},
		Timestamp: time.Now().Format(time.RFC3339), // Discord wants ISO8601; RFC3339 is an extension of ISO8601 and should be completely compatible.
		Title:     animeProfile.Username,
	}

	return embed
}
