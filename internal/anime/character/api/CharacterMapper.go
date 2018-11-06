package api

import (
	"drdgvhbh/discordbot/internal/anime/character/entity"
)

type CharacterMapper interface {
	MapTo(*Character) *entity.Character
}
