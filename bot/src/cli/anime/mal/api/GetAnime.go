package api

import (
	"Drd-Discord-Bot/bot/src/cli/anime/mal/api/response"
	"encoding/json"
	"fmt"

	"github.com/parnurzeal/gorequest"
)

func GetAnime(
	animeID int,
) (*response.AnimeResponse, error) {
	var animeURL = fmt.Sprintf("https://api.jikan.moe/v3/anime/%d/", animeID)
	_, animeResponseBody, errs := gorequest.New().Post(animeURL).End()
	if errs != nil {
		return nil, fmt.Errorf(
			"Failed to retrieve MAL Anime with id: %d", animeID)
	}

	animeResponse := &response.AnimeResponse{}
	json.Unmarshal([]byte(animeResponseBody), animeResponse)

	return animeResponse, nil
}
