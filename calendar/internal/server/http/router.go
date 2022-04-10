package http

import (
	"github.com/gorilla/mux"
	"http/internal/handlers"
	"http/internal/middleware/auth"
	"net/http"
)

func (s *Server) Register(r *mux.Router) {
	handler := handlers.Handler{}
	r.HandleFunc("/login", handler.Login).Methods("POST")
	r.HandleFunc("/logout", handler.Logout).Methods("GET")

	api := r.PathPrefix("/api").Subrouter()
	api.Use(auth.CheckTokenMiddleware)
	api.Path("/users").Handler(http.HandlerFunc(handler.UpdateUserHandler)).Methods("PUT")
	api.Path("/events").Handler(http.HandlerFunc(handler.GetEventsHandler)).Methods("GET")
	api.Path("/events").Handler(http.HandlerFunc(handler.CreateEventHandler)).Methods("POST")
	api.Path("/event/{id}").Handler(http.HandlerFunc(handler.GetEventByIdHandler)).Methods("GET")
	api.Path("/event/{id}").Handler(http.HandlerFunc(handler.UpdateEventHandler)).Methods("PUT")

	http.Handle("/", r)
}
