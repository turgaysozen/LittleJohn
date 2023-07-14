package dto

type Portfolio struct {
	User   string  `json:"user"`
	Stocks []Stock `json:"stocks"`
}

type Stock struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

type StockHistory struct {
	Symbol   string   `json:"symbol"`
	Prices   []Price  `json:"prices"`
	Page     int      `json:"page"`
	PageSize int      `json:"pageSize"`
	Total    int      `json:"total"`
}

type Price struct {
	Date  string `json:"date"`
	Price string `json:"price"`
}