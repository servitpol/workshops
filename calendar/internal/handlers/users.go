package handlers

import (
	"encoding/json"
	"http/internal/repository"
	"net/http"
)

type User struct {
	id       int
	Username string
	Password string
	Email    string
	Token    string
}

func (h Handler) Login(writer http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var request map[string]string
	err := decoder.Decode(&request)
	if err != nil {
		return
	}

	params := &User{
		Username: request["username"],
		Password: request["password"],
	}

	storage := repository.NewStorage(h.Storage)
	err = storage.Db.Login(params)
	if err != nil {
		return
	}
	writer.WriteHeader(200)
	writer.Write([]byte("This is LoginHandler"))
}

func (h Handler) Logout(writer http.ResponseWriter, r *http.Request) {
	writer.WriteHeader(200)
	writer.Write([]byte("This is LogoutHandler"))
}

func (h Handler) UpdateUserHandler(writer http.ResponseWriter, r *http.Request) {
	writer.WriteHeader(200)
	writer.Write([]byte("This is UpdateUserHandler"))
}
