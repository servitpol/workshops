package handlers

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"http/internal/config"
	"http/internal/middleware/auth"
	"http/internal/models"
	"http/internal/repository"
	"net/http"
	"time"
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

	storage := repository.NewStorage(h.Storage)
	user, _ := storage.Db.GetUserByUsername(payload.Username)

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

	cfg := config.GetConfig()

	token, _ := jwt.ParseWithClaims(
		clientToken,
		&auth.JwtClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(cfg.Jwt.Secret), nil
		},
	)
	claims, _ := token.Claims.(*auth.JwtClaim)
	claims.ExpiresAt = time.Now().Local().Unix() - 100000

	writer.WriteHeader(200)
	writer.Write([]byte("successful loged out"))
}

func (h Handler) UpdateUserHandler(writer http.ResponseWriter, r *http.Request) {
	writer.WriteHeader(200)
	writer.Write([]byte("This is UpdateUserHandler"))
}
