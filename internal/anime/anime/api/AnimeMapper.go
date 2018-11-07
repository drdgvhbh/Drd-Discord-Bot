package api

import (
	"drdgvhbh/discordbot/internal/anime/anime/entity"
)

type AnimeMapper interface {
	MapTo(*Anime) *entity.Anime
}
