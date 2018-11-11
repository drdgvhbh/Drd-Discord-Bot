package api

import (
	"drdgvhbh/discordbot/internal/entity"
)

type AnimeMapper interface {
	MapTo(*Anime) *entity.Anime
}
