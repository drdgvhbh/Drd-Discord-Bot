package repository

import (
	"drdgvhbh/discordbot/internal/entity"
)

type CharacterRepository interface {
	SearchCharactersByName(
		name string) ([]*entity.AnimeCharacter, error)
}
