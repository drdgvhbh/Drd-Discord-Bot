package api

import (
	"time"
)

type Stock struct {
	ID          string
	Price       float64
	LastUpdated time.Time
	ImageURL    string
}
