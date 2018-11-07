package entity

type Anime struct {
	id    string
	title string
}

func (anime Anime) ID() string {
	return anime.id
}

func (anime Anime) Title() string {
	return anime.title
}

func CreateAnime(id string, title string) *Anime {
	return &Anime{
		id:    id,
		title: title,
	}
}
