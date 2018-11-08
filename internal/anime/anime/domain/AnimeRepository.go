package domain

import (
	"drdgvhbh/discordbot/internal/anime/anime/entity"
)

type AnimeRepository interface {
	SearchAnimesByTitle(
		name string) ([]*entity.Anime, error)
}
