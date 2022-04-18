package http

import (
	"github.com/gorilla/mux"
	"http/internal/middleware/auth"
	"net/http"
)

func (s *Server) Register(r *mux.Router) {
	handler := s.Handler
	r.HandleFunc("/login", handler.Login).Methods("POST")
	r.HandleFunc("/logout", handler.Logout).Methods("GET")

	api := r.PathPrefix("/api").Subrouter()
	api.Use(auth.CheckTokenMiddleware)
	api.Path("/users").Handler(http.HandlerFunc(handler.UpdateUser)).Methods("PUT")
	api.Path("/events").Handler(http.HandlerFunc(handler.GetEvents)).Methods("GET")
	api.Path("/events").Handler(http.HandlerFunc(handler.CreateEvent)).Methods("POST")
	api.Path("/event/{id}").Handler(http.HandlerFunc(handler.GetEventById)).Methods("GET")
	api.Path("/event/{id}").Handler(http.HandlerFunc(handler.UpdateEvent)).Methods("PUT")

	http.Handle("/", r)
}
