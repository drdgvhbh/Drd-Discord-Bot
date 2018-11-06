package domain

import (
	"drdgvhbh/discordbot/internal/anime/character/entity"
)

type CharacterRepository interface {
	SearchCharactersByName(
		name string) ([]*entity.Character, error)
}
