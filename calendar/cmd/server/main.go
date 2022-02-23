package main

import (
	"github.com/gorilla/mux"
	"http/internal/server/http"
	"http/internal/services/validator"
)

func main() {
	router := mux.NewRouter()
	server := http.NewServer(&validator.Service{})
	server.Register(router)
	http.StartServer(router)
}
