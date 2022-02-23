package main

import (
	"github.com/gorilla/mux"
	httpServer "http/internal/server/http"
	"http/internal/services/validator"
)

func main() {
	router := mux.NewRouter()
	server := httpServer.NewServer(&validator.Service{})
	server.Register(router)
	httpServer.StartServer(router)
}
