package api

type Character struct {
	MalID            int           `json:"mal_id"`
	URL              string        `json:"url"`
	ImageURL         string        `json:"image_url"`
	Name             string        `json:"name"`
	AlternativeNames []interface{} `json:"alternative_names"`
	Anime            []struct {
		MalID int    `json:"mal_id"`
		Type  string `json:"type"`
		Name  string `json:"name"`
		URL   string `json:"url"`
	} `json:"anime"`
	Manga []interface{} `json:"manga"`
}
