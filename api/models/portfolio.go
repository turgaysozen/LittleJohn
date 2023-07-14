package models

import "math/rand"

type Portfolio struct {
	User   string  `json:"user"`
	Stocks []Stock `json:"stocks"`
}

type Stock struct {
	Symbol  string  `json:"symbol"`
	Price   string  `json:"price"`
	History []Price `json:"history"`
}

type Price struct {
	Date  string `json:"date"`
	Price string `json:"price"`
}

func (s Stock) GetPrice() float64 {
	// Get the current price of the stock (dummy implementation)
	return rand.Float64() * 100
}
