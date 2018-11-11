package entity

type Anime struct {
	id       string
	title    string
	score    float64
	imageURL string
}

func (anime Anime) ID() string {
	return anime.id
}

func (anime Anime) Title() string {
	return anime.title
}

func (anime Anime) Score() float64 {
	return anime.score
}

func (anime Anime) ImageURL() string {
	return anime.imageURL
}

type NewAnimeParams struct {
	ID       string
	Title    string
	Score    float64
	ImageURL string
}

func NewAnime(
	params *NewAnimeParams,
) *Anime {
	return &Anime{
		id:       params.ID,
		title:    params.Title,
		score:    params.Score,
		imageURL: params.ImageURL,
	}
}
