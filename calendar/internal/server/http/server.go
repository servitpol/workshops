package http

import (
	"github.com/gorilla/mux"
	"http/internal/config"
	"log"
	"net/http"
	"time"
)

type Validator interface {
	Validate(interface{}) error
}

type Server struct {
	Valid Validator
}

func NewServer(valid Validator) *Server {
	return &Server{Valid: valid}
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
