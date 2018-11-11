package repository

import (
	"drdgvhbh/discordbot/internal/entity"
)

type AnimeRepository interface {
	SearchAnimesByTitle(
		name string) ([]*entity.Anime, error)
}
