package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/turgaysozen/littlejohn/api/handlers"
	"github.com/turgaysozen/littlejohn/api/middlewares"
	"github.com/turgaysozen/littlejohn/api/models"
)

type Server struct {
	router *mux.Router
}

func NewServer() *Server {
	return &Server{
		router: mux.NewRouter(),
	}
}

var GeneratedStocks map[string][]models.Stock

func (s *Server) InitializeRoutes() {
	// use AuthenticationMiddleware to authenticate users
	s.router.Use(middlewares.AuthenticationMiddleware)

	s.router.HandleFunc("/tickers", handlers.GetPortfolio).Methods("GET")
	s.router.HandleFunc("/tickers/{symbol}/history", handlers.GetStockHistory).Methods("GET")
}

func (s *Server) Start(addr string) {
	log.Printf("Server listening on: %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, s.router))
}
