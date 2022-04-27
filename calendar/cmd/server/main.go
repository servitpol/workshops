package main

import (
	"calendar/internal/repository"
	storage "calendar/internal/repository/postgre"
	"calendar/internal/server/http"
	"calendar/internal/services/validator"
	"github.com/gorilla/mux"
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
