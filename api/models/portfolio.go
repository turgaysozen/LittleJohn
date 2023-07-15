package models

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
