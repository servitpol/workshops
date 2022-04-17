package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"http/internal/middleware/auth"
	"http/internal/models"
	"http/internal/repository"
	"log"
	"net/http"
)

func (h Handler) GetEventById(writer http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	clientToken := auth.GetToken(writer, r)

	storage := repository.NewStorage(h.Storage)
	user, err := storage.Db.GetUserByToken(clientToken)
	e, err := storage.Db.GetEventById(vars["id"])

	df, tf, err := e.TimestampToDateTime(e.TimestampFrom, user.Timezone)
	dt, tt, err := e.TimestampToDateTime(e.TimestampTo, user.Timezone)

	event := models.EventResult{
		Id:       user.Id,
		Title:    e.Title,
		DateFrom: df,
		DateTo:   dt,
		TimeFrom: tf,
		TimeTo:   tt,
	}

	js, err := json.Marshal(event)
	if err != nil {
		writer.WriteHeader(403)
		writer.Write([]byte(err.Error()))
		return
	}
	writer.WriteHeader(200)
	writer.Write(js)
}

func (h Handler) CreateEvent(writer http.ResponseWriter, r *http.Request) {

	var e models.Event
	var request map[string]string

	clientToken := auth.GetToken(writer, r)

	storage := repository.NewStorage(h.Storage)
	user, _ := storage.Db.GetUserByToken(clientToken)

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		log.Println(err)
	}

	tf, err := e.TimeToTimestamp(request["time_from"], user.Timezone)
	if err != nil {
		log.Println(err)
	}
	tt, err := e.TimeToTimestamp(request["time_to"], user.Timezone)
	if err != nil {
		log.Println(err)
	}

	payload := models.Event{
		Uid:           user.Id,
		Title:         request["title"],
		Description:   request["description"],
		TimestampFrom: tf,
		TimestampTo:   tt,
	}

	id, err := storage.Db.CreateEvent(payload)

	js, err := json.Marshal(id)
	writer.WriteHeader(200)
	writer.Write(js)
}

func (h Handler) UpdateEventHandler(writer http.ResponseWriter, r *http.Request) {
	writer.WriteHeader(200)
	writer.Write([]byte("This is UpdateEventHandler"))
}

func (h Handler) GetEvents(writer http.ResponseWriter, r *http.Request) {
	writer.WriteHeader(200)
	writer.Write([]byte("This is GetEventsHandler"))
}
