package api

import (
	"drdgvhbh/discordbot/internal/entity"
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

var characterRepository *CharacterRepository

func ProvideCharacterRepository(
	mapper CharacterMapper,
) *CharacterRepository {
	if characterRepository != nil {
		return characterRepository
	}

	characterRepository = CreateCharacterRepository(mapper)

	return characterRepository
}

func (
	characterRepository CharacterRepository,
) SearchCharactersByName(
	name string,
) ([]*entity.AnimeCharacter, error) {
	var endpoint = fmt.Sprintf(
		"https://api.jikan.moe/v3/search/character/?q=%s&page=1",
		name)

	response, searchCharacterResponseBody, errs := gorequest.New().Get(endpoint).End()
	if response.StatusCode != 200 || errs != nil {
		return nil, fmt.Errorf(
			"Failed to find any anime character with name: %s", name)
	}

	searchCharacterResponse := &SearchCharacterResponse{}
	json.Unmarshal([]byte(searchCharacterResponseBody), searchCharacterResponse)

	characters := funk.Map(
		searchCharacterResponse.Results,
		func(x *Character) *entity.AnimeCharacter {
			return characterRepository.mapper.MapTo(x)
		}).([]*entity.AnimeCharacter)
	return characters, nil
}
