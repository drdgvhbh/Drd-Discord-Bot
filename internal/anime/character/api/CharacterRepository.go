package api

import (
	"drdgvhbh/discordbot/internal/anime/character/entity"
	"encoding/json"
	"fmt"

	"github.com/parnurzeal/gorequest"
	"github.com/thoas/go-funk"
)

type CharacterRepository struct {
	mapper CharacterMapper
}

func CreateCharacterRepository(
	mapper CharacterMapper,
) *CharacterRepository {
	return &CharacterRepository{
		mapper,
	}
}

func (
	characterRepository CharacterRepository,
) SearchCharactersByName(
	name string,
) ([]*entity.Character, error) {
	var endpoint = fmt.Sprintf(
		"https://api.jikan.moe/v3/search/character/?q=%s&page=1",
		name)

	_, searchCharacterResponseBody, errs := gorequest.New().Get(endpoint).End()
	if errs != nil {
		return nil, fmt.Errorf(
			"Failed to find any anime character with name: %s", name)
	}

	searchCharacterResponse := &SearchCharacterResponse{}
	json.Unmarshal([]byte(searchCharacterResponseBody), searchCharacterResponse)

	characters := funk.Map(
		searchCharacterResponse.Results,
		func(x *Character) *entity.Character {
			return characterRepository.mapper.MapTo(x)
		}).([]*entity.Character)
	return characters, nil
}
