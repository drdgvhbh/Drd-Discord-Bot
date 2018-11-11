package mapper

import (
	"drdgvhbh/discordbot/internal/anime/anime/api"
	"drdgvhbh/discordbot/internal/entity"
	"strconv"
)

type AnimeMapper struct {
}

func CreateAnimeMapper() *AnimeMapper {
	return &AnimeMapper{}
}

var animeMapper *AnimeMapper

func ProvideAnimeMapper() *AnimeMapper {
	if animeMapper != nil {
		return animeMapper
	}
	animeMapper = CreateAnimeMapper()

	return animeMapper
}

func (animeMapper AnimeMapper) MapTo(
	anime *api.Anime,
) *entity.Anime {
	return entity.NewAnime(
		&entity.NewAnimeParams{
			ID:       strconv.Itoa(anime.MalID),
			Title:    anime.Title,
			Score:    anime.Score,
			ImageURL: anime.ImageURL,
		},
	)
}
