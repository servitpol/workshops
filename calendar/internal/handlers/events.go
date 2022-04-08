package handlers

import (
	"http/internal/repository"
	"net/http"
)

type Event struct {
	id          int
	Uid         int
	Title       string
	Description string
	Timezone    string
	Duration    int
	EventTime   int
}

func (h Handler) GetEventsHandler(writer http.ResponseWriter, r *http.Request) {
	storage := repository.NewStorage(h.Storage)
	err := storage.Db.GetEvents()
	if err != nil {
		return
	}

	writer.WriteHeader(200)
	writer.Write([]byte("This is GetEventsHandler"))
}

func (h Handler) CreateEventHandler(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(200)
	writer.Write([]byte("This is CreateEventHandler"))
}

func (h Handler) GetEventByIdHandler(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(200)
	writer.Write([]byte("This is GetEventByIdHandler"))
}

func (h Handler) UpdateEventHandler(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(200)
	writer.Write([]byte("This is UpdateEventHandler"))
}
