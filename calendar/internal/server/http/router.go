package http

import (
	"github.com/gorilla/mux"
	"http/internal/handlers"
	"net/http"
)

func (s *Server) Register(r *mux.Router) {
	handler := handlers.Handler{}
	r.HandleFunc("/login", handler.Login).Methods("POST")
	r.HandleFunc("/logout", handler.Logout).Methods("GET")
	r.HandleFunc("/api/user", handler.UpdateUserHandler).Methods("PUT")

	r.HandleFunc("/api/events", handler.GetEventsHandler).Methods("GET")
	r.HandleFunc("/api/events", handler.CreateEventHandler).Methods("POST")
	r.HandleFunc("/api/event/{id}", handler.GetEventByIdHandler).Methods("GET")
	r.HandleFunc("/api/event/{id}", handler.UpdateEventHandler).Methods("PUT")

	http.Handle("/", r)
}
