package services

import (
	"math/rand"
	"time"

	"github.com/turgaysozen/littlejohn/api/models"
	"github.com/turgaysozen/littlejohn/dto"
	"github.com/turgaysozen/littlejohn/dummy_data"
	logger "github.com/turgaysozen/littlejohn/utils"
)

var generatedStocks = make(map[string][]models.Stock)

func init() {
	generateStocks()
}

func generateStocks() {
	for _, username := range dummy_data.ValidUsernames {
		userStocks := make([]models.Stock, 0)

		// Randomly select stocks for the user
		for i := 0; i < rand.Intn(10)+1; i++ { // Generate a random portfolio size between 1 and 10
			stockIndex := rand.Intn(len(dummy_data.Stocks))
			stock := dummy_data.Stocks[stockIndex]

			// Check if the stock is already in the user's portfolio
			stockExists := false
			for _, existingStock := range userStocks {
				if existingStock.Symbol == stock.Symbol {
					stockExists = true
					break
				}
			}

			// If the stock doesn't exist in the user's portfolio, add it
			if !stockExists {
				userStocks = append(userStocks, stock)
			}
		}

		logger.Info.Println("Stocks are generated as initial for user:", username)
		generatedStocks[username] = userStocks
	}
}

func GetPortfolioByUsername(username string) dto.Portfolio {
	stocks, found := generatedStocks[username]
	if !found {
		logger.Error.Println("Portfolio cannot find for user:", username)
		return dto.Portfolio{
			User:   username,
			Stocks: []dto.Stock{},
		}
	}

	// Convert the stocks from models.Stock to dto.Stock
	dtoStocks := make([]dto.Stock, len(stocks))
	for i, stock := range stocks {
		dtoStocks[i] = dto.Stock{
			Symbol: stock.Symbol,
			Price:  stock.Price,
		}
	}

	return dto.Portfolio{
		User:   username,
		Stocks: dtoStocks,
	}
}

func GetStockBySymbol(symbol string) *models.Stock {
	for _, stock := range dummy_data.Stocks {
		if stock.Symbol == symbol {
			return &stock
		}
	}
	return nil
}

func GetStockHistoryBySymbol(symbol string, page, pageSize int) (dto.StockHistory, bool) {
	// Retrieve the stock by symbol from the dummy data
	stock := GetStockBySymbol(symbol)
	if stock == nil {
		logger.Error.Println("Stock history cannot find for symbol:", symbol)
		return dto.StockHistory{
			Symbol:   symbol,
			Prices:   []dto.Price{},
			Page:     page,
			PageSize: pageSize,
			Total:    0,
		}, false
	}

	// Get the total number of historical prices for the stock
	total := len(stock.History)

	// Calculate the start and end indices for the requested page
	startIndex := (page - 1) * pageSize
	endIndex := startIndex + pageSize

	// Adjust the end index if it exceeds the total number of prices
	if endIndex > total {
		endIndex = total
	}

	// Get the prices for the requested page
	prices := stock.History[startIndex:endIndex]

	return dto.StockHistory{
		Symbol:   symbol,
		Prices:   convertPrices(prices),
		Page:     page,
		PageSize: pageSize,
		Total:    total,
	}, true
}

func convertPrices(prices []models.Price) []dto.Price {
	// Get the current date
	currentDate := time.Now()

	// Calculate the date 90 days ago
	maxDays := 90
	maxDate := currentDate.AddDate(0, 0, -maxDays)

	// Filter out prices older than the last 90 days
	filteredPrices := make([]models.Price, 0)
	for _, price := range prices {
		priceDate, err := time.Parse("2006-01-02", price.Date)
		if err == nil && priceDate.After(maxDate) {
			filteredPrices = append(filteredPrices, price)
		}
	}

	// Convert the filtered prices to DTO format
	converted := make([]dto.Price, len(filteredPrices))
	for i, price := range filteredPrices {
		converted[i] = dto.Price{
			Date:  price.Date,
			Price: price.Price,
		}
	}
	return converted
}
