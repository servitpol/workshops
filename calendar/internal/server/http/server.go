package http

import (
	"github.com/gorilla/mux"
	"http/internal/config"
	"http/internal/handlers"
	"http/internal/repository"
	"log"
	"net/http"
	"time"
)

type Validator interface {
	Validate(interface{}) error
}

type Server struct {
	Valid    Validator
	Storage  repository.Storage
	Handlers handlers.Handler
}

func NewServer(valid Validator, s repository.Storage) *Server {
	return &Server{
		Valid:    valid,
		Storage:  s,
		Handlers: handlers.Handler{},
	}
}

func StartServer(router *mux.Router) {
	srv := &http.Server{
		Handler:      router,
		Addr:         getAddr(),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func getAddr() string {
	cfg := config.GetConfig()
	addr := cfg.Lsn.BindIP + ":" + cfg.Lsn.Port
	return addr
}
