package handlers

import (
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"http/internal/config"
	"http/internal/middleware/auth"
	"http/internal/repository"
	"net/http"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type LoginPayload struct {
	Password string `json:"password"`
	Username string `json:"email"`
}

func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}

	return nil
}

func (h Handler) Login(writer http.ResponseWriter, r *http.Request) {
	var user *User
	var payload LoginPayload

	//need to add validate login|password fields
	decoder := json.NewDecoder(r.Body)
	var request map[string]string
	err := decoder.Decode(&request)
	if err != nil {
		return
	}

	payload = LoginPayload{
		Password: request["password"],
		Username: request["username"],
	}

	storage := repository.NewStorage(h.Storage)
	storagePassword, err := storage.Db.GetUser(payload.Username)

	user = &User{
		Password: storagePassword,
	}
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
	tokenResponse := LoginResponse{
		Token: signedToken,
	}

	js, err := json.Marshal(tokenResponse)
	if err != nil {
		writer.WriteHeader(403)
		writer.Write([]byte("rasas"))
		return
	}
	writer.WriteHeader(200)
	writer.Write(js)
}

func (h Handler) Logout(writer http.ResponseWriter, r *http.Request) {
	writer.WriteHeader(200)
	writer.Write([]byte("This is LogoutHandler"))
}

func (h Handler) UpdateUserHandler(writer http.ResponseWriter, r *http.Request) {
	writer.WriteHeader(200)
	writer.Write([]byte("This is UpdateUserHandler"))
}
