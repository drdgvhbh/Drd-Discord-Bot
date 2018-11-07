package entity

type Anime struct {
	id    string
	title string
	score float64
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

func CreateAnime(id string, title string, score float64) *Anime {
	return &Anime{
		id:    id,
		title: title,
		score: score,
	}
}
