package handlers

import (
	"calendar/internal/config"
	"calendar/internal/middleware/auth"
	"calendar/internal/models"
	"encoding/json"
	"log"
	"net/http"
)

func (h Handler) Login(writer http.ResponseWriter, r *http.Request) {

	var payload models.LoginPayload

	decoder := json.NewDecoder(r.Body)
	var request map[string]string
	err := decoder.Decode(&request)
	if err != nil {
		return
	}

	payload = models.LoginPayload{
		Password: request["password"],
		Username: request["username"],
	}

	user, _ := h.Storage.Db.GetUserByUsername(payload.Username)

	if err != nil {
		writer.WriteHeader(403)
		writer.Write([]byte("resource not found"))
		return
	}

	err = user.CheckPassword(payload.Password)
	if err != nil {
		writer.WriteHeader(403)
		writer.Write([]byte("wrong password"))
		return
	}

	cfg := config.GetConfig()
	jwtWrapper := auth.JwtWrapper{
		SecretKey:       cfg.Jwt.Secret,
		Issuer:          cfg.Jwt.Issuer,
		ExpirationHours: cfg.Jwt.ExHours,
	}
	signedToken, _ := jwtWrapper.GenerateToken(payload.Username)
	tokenResponse := models.LoginResponse{
		Token: signedToken,
	}

	err = h.Storage.Db.UpdateUserToken(signedToken, user.Id)
	if err != nil {
		log.Println(err)
	}

	js, err := json.Marshal(tokenResponse)
	if err != nil {
		writer.WriteHeader(403)
		writer.Write([]byte(err.Error()))
		return
	}
	writer.WriteHeader(200)
	writer.Write(js)
}

func (h Handler) Logout(writer http.ResponseWriter, r *http.Request) {
	clientToken := auth.GetToken(writer, r)

	user, err := h.Storage.Db.GetUserByToken(clientToken)
	err = h.Storage.Db.UpdateUserToken("", user.Id)
	if err != nil {
		log.Println(err)
	}

	writer.WriteHeader(200)
	writer.Write([]byte("successful loged out"))
}

func (h Handler) UpdateUser(writer http.ResponseWriter, r *http.Request) {

	var request map[string]string
	clientToken := auth.GetToken(writer, r)

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		log.Println(err)
	}
	err = h.Storage.Db.UpdateUserTimezone(clientToken, request["timezone"])
	if err != nil {
		log.Println(err)
	}

	writer.WriteHeader(200)
	writer.Write([]byte("Successfully saved"))
}
