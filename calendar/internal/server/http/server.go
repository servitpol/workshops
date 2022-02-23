package http

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

type Validator interface {
	Validate(interface{}) error
}

type Server struct {
	valid Validator
}

func NewServer(valid Validator) *Server {
	return &Server{valid: valid}
}

func StartServer(router *mux.Router) {
	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
