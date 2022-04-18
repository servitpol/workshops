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
	Valid   Validator
	Handler handlers.Handler
}

func NewServer(valid Validator, r repository.Db) *Server {
	return &Server{
		Valid: valid,
		Handler: struct {
			Storage repository.Db
		}{Storage: r},
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
