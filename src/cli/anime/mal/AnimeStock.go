package mal

import (
	"cli/anime/mal/api"
	"fmt"
	"math"
)

type animeStock struct {
	name        string
	imageURL    string
	marketPrice float32
	characterID int
	animeID     int
}

type AnimeStock interface {
	Name() string
	ImageURL() string
	MarketPrice() float32
}

func (animeStock animeStock) Name() string {
	return animeStock.name
}

func (animeStock animeStock) ImageURL() string {
	return animeStock.imageURL
}

func (animeStock animeStock) MarketPrice() float32 {
	return animeStock.marketPrice
}

func CreateAnimeStock(
	characterID int,
	animeID int,
) (AnimeStock, error) {
	animeResponse, err := api.GetAnime(animeID)
	if err != nil {
		return nil, fmt.Errorf(
			fmt.Sprintf(
				"Failed to create anime stock. Could not retrieve anime with id: %d",
				animeID))
	}

	animeCharacter, err := api.GetAnimeCharacter(characterID)
	if err != nil {
		return nil, fmt.Errorf(
			fmt.Sprintf(
				"Failed to create anime stock. Could not retrieve character with id: %d",
				characterID))
	}

	favoritesPriceBonus := math.Log(
		float64(animeCharacter.MemberFavorites)) / math.Log(10000.0)
	marketPrice := animeResponse.Score + favoritesPriceBonus

	animeStock := &animeStock{
		characterID: characterID,
		animeID:     animeID,
		marketPrice: float32(marketPrice),
		name: fmt.Sprintf(
			"%s-%s",
			animeCharacter.Name,
			animeResponse.TitleEnglish),
		imageURL: animeCharacter.ImageURL,
	}
	return animeStock, nil
}
