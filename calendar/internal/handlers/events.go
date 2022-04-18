package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"http/internal/middleware/auth"
	"http/internal/models"
	"log"
	"net/http"
	"strconv"
)

func (h Handler) GetEvents(writer http.ResponseWriter, r *http.Request) {

	clientToken := auth.GetToken(writer, r)

	user, _ := h.Storage.Db.GetUserByToken(clientToken)
	events, err := h.Storage.Db.GetEvents()
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

	user, err := h.Storage.Db.GetUserByToken(clientToken)

	e, err := h.Storage.Db.GetEventById(vars["id"])
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

	user, _ := h.Storage.Db.GetUserByToken(clientToken)

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

	id, err := h.Storage.Db.CreateEvent(payload)

	js, err := json.Marshal(id)
	writer.WriteHeader(200)
	writer.Write(js)
}

func (h Handler) UpdateEvent(writer http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	var e models.Event
	var request map[string]string

	clientToken := auth.GetToken(writer, r)

	user, _ := h.Storage.Db.GetUserByToken(clientToken)

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

	eventId, err := strconv.Atoi(vars["id"])

	payload := models.Event{
		Uid:           user.Id,
		Title:         request["title"],
		Description:   request["description"],
		TimestampFrom: tf,
		TimestampTo:   tt,
	}

	err = h.Storage.Db.UpdateEvent(payload, eventId)

	writer.WriteHeader(200)
	writer.Write([]byte("Successfully saved"))
}
