package main

import (
	"github.com/gorilla/mux"
	"http/internal/repository"
	storage "http/internal/repository/postgre"
	"http/internal/server/http"
	"http/internal/services/validator"
)

func main() {
	pgConn := storage.New()
	s := &storage.Repository{Pool: pgConn}

	str := *repository.NewStorage(s)
	server := http.NewServer(&validator.Service{}, str)

	router := mux.NewRouter()
	server.Register(router)
	http.StartServer(router)
}
