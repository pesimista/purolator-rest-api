package app

import (
	"github.com/pesimista/purolator-rest-api/internal/api"
)

func Run() {
	server := api.NewServer()
	// server.CreateServer()

	server.SetRoutes()
	server.Run()
}
