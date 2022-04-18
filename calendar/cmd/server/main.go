package main

import (
	"github.com/gorilla/mux"
	"http/internal/repository"
	"http/internal/repository/postgre"
	"http/internal/server/http"
	"http/internal/services/validator"
)

func main() {
	router := mux.NewRouter()
	storage := &postgre.Postgres{}
	str := *repository.NewStorage(storage)
	server := http.NewServer(&validator.Service{}, str)
	server.Register(router)
	http.StartServer(router)
}
