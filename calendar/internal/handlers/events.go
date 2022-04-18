package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"http/internal/middleware/auth"
	"http/internal/models"
	"log"
	"net/http"
)

func (h Handler) GetEvents(writer http.ResponseWriter, r *http.Request) {

	clientToken := auth.GetToken(writer, r)
	user, _ := h.R.Db.GetUserByToken(clientToken)
	events, err := h.R.Db.GetEvents()
	if err != nil {
		log.Println(err)
	}

	var mSlice = make([]models.EventResult, 0)
	for _, event := range events {
		e := event.MakeApiData(user)
		mSlice = append(mSlice, e)
	}

	js, err := json.Marshal(mSlice)
	writer.WriteHeader(200)
	writer.Write(js)
}

func (h Handler) GetEventById(writer http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	clientToken := auth.GetToken(writer, r)
	user, err := h.R.Db.GetUserByToken(clientToken)

	e, err := h.R.Db.GetEventById(vars["id"])
	event := e.MakeApiData(user)

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

	user, _ := h.R.Db.GetUserByToken(clientToken)

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

	id, err := h.R.Db.CreateEvent(payload)

	js, err := json.Marshal(id)
	writer.WriteHeader(200)
	writer.Write(js)
}

func (h Handler) UpdateEventHandler(writer http.ResponseWriter, r *http.Request) {
	writer.WriteHeader(200)
	writer.Write([]byte("This is UpdateEventHandler"))
}
