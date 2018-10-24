package mal

import (
	"cli/anime/mal/api"
	"encoding/json"
	"fmt"
	"math"

	"github.com/parnurzeal/gorequest"
)

type AnimeStock struct {
	CharacterID int
	AnimeID     int
	OwnerID     string
}

type Stock interface {
	GetMarketPrice() (float64, error)
}

func (animeStock AnimeStock) GetMarketPrice() (float64, error) {
	var animeURL = fmt.Sprintf("https://api.jikan.moe/v3/anime/%d/", animeStock.AnimeID)
	_, animeResponseBody, errs := gorequest.New().Post(animeURL).End()
	if errs != nil {
		return -1, fmt.Errorf(
			"Failed to retrieve MAL Anime with id: %d", animeStock.AnimeID)
	}

	animeResponse := &api.AnimeResponse{}
	json.Unmarshal([]byte(animeResponseBody), animeResponse)

	var characterURL = fmt.Sprintf("https://api.jikan.moe/v3/character/%d/", animeStock.CharacterID)
	_, characterResponseBody, errs := gorequest.New().Post(characterURL).End()
	if errs != nil {
		return -1, fmt.Errorf(
			"Failed to retrieve MAL Character with id: %d", animeStock.CharacterID)
	}

	characterResponse := &api.AnimeCharacterResponse{}
	json.Unmarshal([]byte(characterResponseBody), characterResponse)

	price := animeResponse.Score + math.Log(float64(characterResponse.MemberFavorites))/math.Log(10000.0)
	return price, nil
}
