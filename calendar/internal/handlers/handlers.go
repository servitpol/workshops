package handlers

import (
	"http/internal/repository"
	"net/http"
)

type Handlers interface {
	Login(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)

	GetEvents(w http.ResponseWriter, r *http.Request)
	GetEventById(w http.ResponseWriter, r *http.Request)
	CreateEvent(w http.ResponseWriter, r *http.Request)
	UpdateEventHandler(w http.ResponseWriter, r *http.Request)
}

type Handler struct {
	H Handlers
	R repository.Db
}
