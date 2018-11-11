package entity

import "time"

type Stock struct {
	id          string
	price       float64
	lastUpdated time.Time
	imageURL    string
}

func (stock *Stock) ID() string {
	return stock.id
}

func (stock *Stock) Price() float64 {
	return stock.price
}

func (stock *Stock) LastUpdated() time.Time {
	return stock.lastUpdated
}

func (stock *Stock) ImageURL() string {
	return stock.imageURL
}

type NewStockParams struct {
	ID          string
	Price       float64
	LastUpdated time.Time
	ImageURL    string
}

func NewStock(params *NewStockParams) *Stock {
	return &Stock{
		id:          params.ID,
		price:       params.Price,
		imageURL:    params.ImageURL,
		lastUpdated: params.LastUpdated,
	}
}
