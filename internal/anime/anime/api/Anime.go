package api

import "time"

type Anime struct {
	MalID     int       `json:"mal_id"`
	URL       string    `json:"url"`
	ImageURL  string    `json:"image_url"`
	Title     string    `json:"title"`
	Airing    bool      `json:"airing"`
	Synopsis  string    `json:"synopsis"`
	Type      string    `json:"type"`
	Episodes  int       `json:"episodes"`
	Score     float64   `json:"score"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	Members   int       `json:"members"`
	Rated     string    `json:"rated"`
}
