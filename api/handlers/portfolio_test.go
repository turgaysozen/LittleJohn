package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/turgaysozen/littlejohn/dto"
)

func TestGetPortfolio(t *testing.T) {
	// Create a new HTTP request to test the GetPortfolio function
	req, err := http.NewRequest("GET", "/tickers", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Set the context value for the authenticated username
	ctx := req.Context()
	ctx = context.WithValue(ctx, "username", "user1")
	req = req.WithContext(ctx)

	// Create a response recorder to capture the response
	rr := httptest.NewRecorder()

	// Call the GetPortfolio function with the test request and response recorder
	GetPortfolio(rr, req)

	// Check the response status code
	if rr.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, rr.Code)
	}

	var response []dto.Stock

	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	// Check the return type
	if len(response) == 0 {
		t.Errorf("expected stocks to be present in the response")
	}
}

func TestGetStockHistory(t *testing.T) {
	// Create a request with query parameters
	req := httptest.NewRequest("GET", "/tickers/V/history?page=1&pageSize=1", nil)

	// Create a response recorder to record the response
	res := httptest.NewRecorder()

	// Set the stock symbol path parameter
	req = mux.SetURLVars(req, map[string]string{"symbol": "V"})

	// Call the handler function
	GetStockHistory(res, req)

	// Check the response status code
	if res.Code != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, res.Code)
	}
}
