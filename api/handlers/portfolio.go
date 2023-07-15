package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/turgaysozen/littlejohn/api/services"
	logger "github.com/turgaysozen/littlejohn/utils"
)

func GetPortfolio(w http.ResponseWriter, r *http.Request) {
	// Get the authenticated username from the request context
	username := r.Context().Value("username").(string)

	// Retrieve the user's portfolio from the dummy data
	portfolio := services.GetPortfolioByUsername(username)
	logger.Info.Println("Getting portfolio for user:", username, "portfolio:", portfolio)

	// Return the stocks in the user's portfolio as the response
	jsonResponse, err := json.Marshal(portfolio.Stocks)
	if err != nil {
		logger.Error.Println("An error occurred while marshalling stocks of portfolio user, err:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func GetStockHistory(w http.ResponseWriter, r *http.Request) {
	// Get the stock symbol from the request path parameters
	params := mux.Vars(r)
	symbol := params["symbol"]

	pageStr := r.URL.Query().Get("page")
	if pageStr == "" {
		// Set default value for page if not provided
		pageStr = "1"
	}

	pageSizeStr := r.URL.Query().Get("pageSize")
	if pageSizeStr == "" {
		// Set default value for pageSize if not provided
		pageSizeStr = "10"
	}

	// Parse the page and pageSize values to integers
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		logger.Error.Println("Invalid page parameter, err:", err)
		http.Error(w, "Invalid page parameter", http.StatusBadRequest)
		return
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		logger.Error.Println("Invalid pageSize parameter, err:", err)
		http.Error(w, "Invalid pageSize parameter", http.StatusBadRequest)
		return
	}

	// Retrieve the stock history based on the symbol and pagination parameters
	stockHistory, found := services.GetStockHistoryBySymbol(symbol, page, pageSize)
	logger.Info.Println("Getting stock history for symbol:", symbol, "with page:", page, "pageSize:", pageSize)

	if !found {
		logger.Info.Println("Stock cannot find for given symbol:", symbol)
		http.NotFound(w, r)
		return
	}

	// Convert the stock history to JSON
	jsonResponse, err := json.Marshal(stockHistory)
	if err != nil {
		logger.Error.Println("An error occurred while marshalling stocks history, err:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
