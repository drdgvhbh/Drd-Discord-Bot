package api

import (
	"drdgvhbh/discordbot/internal/anime/anime/entity"
	"encoding/json"
	"fmt"

	"github.com/parnurzeal/gorequest"
	"github.com/thoas/go-funk"
)

type AnimeRepository struct {
	mapper AnimeMapper
}

func CreateAnimeRepository(
	mapper AnimeMapper,
) *AnimeRepository {
	return &AnimeRepository{
		mapper,
	}
}

var animeRepository *AnimeRepository

func ProvideAnimeRepository(
	mapper AnimeMapper,
) *AnimeRepository {
	if animeRepository != nil {
		return animeRepository
	}

	animeRepository = CreateAnimeRepository(mapper)

	return animeRepository
}

func (
	animeRepository AnimeRepository,
) SearchAnimesByTitle(
	name string,
) ([]*entity.Anime, error) {
	var endpoint = fmt.Sprintf(
		"https://api.jikan.moe/v3/search/anime/?q=%s&page=1",
		name)

	response, searchAnimeResponseBody, errs := gorequest.New().Get(endpoint).End()
	if response.StatusCode != 200 || errs != nil {
		return nil, fmt.Errorf(
			"Failed to find any anime with name: %s", name)
	}

	searchAnimeResponse := &SearchAnimeResponse{}
	json.Unmarshal([]byte(searchAnimeResponseBody), searchAnimeResponse)

	animes := funk.Map(
		searchAnimeResponse.Results,
		func(x *Anime) *entity.Anime {
			return animeRepository.mapper.MapTo(x)
		}).([]*entity.Anime)

	return animes, nil
}
