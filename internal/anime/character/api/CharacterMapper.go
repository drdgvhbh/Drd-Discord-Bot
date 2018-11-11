package api

import (
	"drdgvhbh/discordbot/internal/entity"
)

type CharacterMapper interface {
	MapTo(*Character) *entity.AnimeCharacter
}
