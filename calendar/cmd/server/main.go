package main

import (
	"github.com/gorilla/mux"
	"http/internal/repository/postgre"
	"http/internal/server/http"
	"http/internal/services/validator"
)

func main() {
	router := mux.NewRouter()
	storage := postgre.New()
	server := http.NewServer(&validator.Service{}, storage)

	server.Register(router)
	http.StartServer(router)
}
