package main

import (
	"github.com/turgaysozen/littlejohn/server"
)

func main() {
	s := server.NewServer()

	s.InitializeRoutes()

	s.Start(":8080")
}
