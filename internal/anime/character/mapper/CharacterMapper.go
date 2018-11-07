package mapper

import (
	"drdgvhbh/discordbot/internal/anime/character/api"
	"drdgvhbh/discordbot/internal/anime/character/entity"
	"strconv"
)

type CharacterMapper struct {
}

func CreateCharacterMapper() *CharacterMapper {
	return &CharacterMapper{}
}

var characterMapper *CharacterMapper

func ProvideCharacterMapper() *CharacterMapper {
	if characterMapper != nil {
		return characterMapper
	}
	characterMapper = CreateCharacterMapper()

	return characterMapper
}

func (characterMapper CharacterMapper) MapTo(
	character *api.Character,
) *entity.Character {
	return entity.CreateCharacter(
		strconv.Itoa(character.MalID),
		character.Name,
	)
}
