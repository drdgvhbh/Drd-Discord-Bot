package api

import (
	"drdgvhbh/discordbot/internal/cli/anime/mal/api/response"
	"encoding/json"
	"fmt"

	"github.com/parnurzeal/gorequest"
)

func GetAnimeCharacter(
	characterID int,
) (*response.AnimeCharacterResponse, error) {
	var characterURL = fmt.Sprintf("https://api.jikan.moe/v3/character/%d/", characterID)
	_, characterResponseBody, errs := gorequest.New().Post(characterURL).End()
	if errs != nil {
		return nil, fmt.Errorf(
			"Failed to retrieve MAL Character with id: %d", characterID)
	}

	characterResponse := &response.AnimeCharacterResponse{}
	json.Unmarshal([]byte(characterResponseBody), characterResponse)

	return characterResponse, nil
}
