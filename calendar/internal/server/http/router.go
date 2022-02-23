package http

import (
	"github.com/gorilla/mux"
	"net/http"
)

func (s *Server) Register(r *mux.Router) {

	r.HandleFunc("/login", s.LoginHandler).
		Methods("POST")

	r.HandleFunc("/logout", s.LogoutHandler).
		Methods("GET")

	r.HandleFunc("/api/user", s.UpdateUserHandler).
		Methods("PUT")

	r.HandleFunc("/api/events", s.GetEventsHandler).
		Methods("GET")

	r.HandleFunc("/api/events", s.CreateEventHandler).
		Methods("POST")

	r.HandleFunc("/api/event/{id}", s.GetEventByIdHandler).
		Methods("GET")

	r.HandleFunc("/api/event/{id}", s.UpdateEventHandler).
		Methods("PUT")

	http.Handle("/", r)
}

func (s *Server) LoginHandler(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(200)
	writer.Write([]byte("This is LoginHandler"))
}

func (s *Server) LogoutHandler(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(200)
	writer.Write([]byte("This is LogoutHandler"))
}

func (s *Server) UpdateUserHandler(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(200)
	writer.Write([]byte("This is UpdateUserHandler"))
}

func (s *Server) GetEventsHandler(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(200)
	writer.Write([]byte("This is GetEventsHandler"))
}

func (s *Server) CreateEventHandler(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(201)
	writer.Write([]byte("This is CreateEventHandler"))
}

func (s *Server) GetEventByIdHandler(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(200)
	writer.Write([]byte("This is GetEventByIdHandler"))
}

func (s *Server) UpdateEventHandler(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(201)
	writer.Write([]byte("This is UpdateEventHandler"))
}
