package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Timezone string `json:"timezone"`
	Token    string `json:"token"`
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
