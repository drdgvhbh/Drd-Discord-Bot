package api

import "time"

type AnimeResponse struct {
	RequestHash        string        `json:"request_hash"`
	RequestCached      bool          `json:"request_cached"`
	RequestCacheExpiry int           `json:"request_cache_expiry"`
	MalID              int           `json:"mal_id"`
	URL                string        `json:"url"`
	ImageURL           string        `json:"image_url"`
	TrailerURL         string        `json:"trailer_url"`
	Title              string        `json:"title"`
	TitleEnglish       string        `json:"title_english"`
	TitleJapanese      string        `json:"title_japanese"`
	TitleSynonyms      []interface{} `json:"title_synonyms"`
	Type               string        `json:"type"`
	Source             string        `json:"source"`
	Episodes           int           `json:"episodes"`
	Status             string        `json:"status"`
	Airing             bool          `json:"airing"`
	Aired              struct {
		From time.Time `json:"from"`
		To   time.Time `json:"to"`
		Prop struct {
			From struct {
				Day   int `json:"day"`
				Month int `json:"month"`
				Year  int `json:"year"`
			} `json:"from"`
			To struct {
				Day   int `json:"day"`
				Month int `json:"month"`
				Year  int `json:"year"`
			} `json:"to"`
		} `json:"prop"`
		String string `json:"string"`
	} `json:"aired"`
	Duration   string      `json:"duration"`
	Rating     string      `json:"rating"`
	Score      float64     `json:"score"`
	ScoredBy   int         `json:"scored_by"`
	Rank       int         `json:"rank"`
	Popularity int         `json:"popularity"`
	Members    int         `json:"members"`
	Favorites  int         `json:"favorites"`
	Synopsis   string      `json:"synopsis"`
	Background interface{} `json:"background"`
	Premiered  string      `json:"premiered"`
	Broadcast  string      `json:"broadcast"`
	Related    struct {
		Adaptation []struct {
			MalID int    `json:"mal_id"`
			Type  string `json:"type"`
			Name  string `json:"name"`
			URL   string `json:"url"`
		} `json:"Adaptation"`
		Prequel []struct {
			MalID int    `json:"mal_id"`
			Type  string `json:"type"`
			Name  string `json:"name"`
			URL   string `json:"url"`
		} `json:"Prequel"`
		Other []struct {
			MalID int    `json:"mal_id"`
			Type  string `json:"type"`
			Name  string `json:"name"`
			URL   string `json:"url"`
		} `json:"Other"`
	} `json:"related"`
	Producers []interface{} `json:"producers"`
	Licensors []struct {
		MalID int    `json:"mal_id"`
		Type  string `json:"type"`
		Name  string `json:"name"`
		URL   string `json:"url"`
	} `json:"licensors"`
	Studios []struct {
		MalID int    `json:"mal_id"`
		Type  string `json:"type"`
		Name  string `json:"name"`
		URL   string `json:"url"`
	} `json:"studios"`
	Genres []struct {
		MalID int    `json:"mal_id"`
		Type  string `json:"type"`
		Name  string `json:"name"`
		URL   string `json:"url"`
	} `json:"genres"`
	OpeningThemes []string `json:"opening_themes"`
	EndingThemes  []string `json:"ending_themes"`
}
