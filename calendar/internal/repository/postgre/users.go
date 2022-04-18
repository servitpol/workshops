package postgre

import (
	"context"
	"http/internal/models"
	"log"
)

func (pg *Postgres) GetUserByUsername(username string) (models.User, error) {

	var u models.User
	sql := "SELECT * FROM users WHERE username=$1"

	mRows, err := pg.db.Query(context.Background(), sql, username)
	if err != nil {
		log.Println(err)
	}

	for mRows.Next() {
		err := mRows.Scan(&u.Id, &u.Username, &u.Password, &u.Email, &u.Timezone, &u.Token)
		if err != nil {
			log.Println(err)
		}
	}

	return u, err
}

func (pg *Postgres) GetUserByToken(token string) (models.User, error) {

	var u models.User
	sql := "SELECT * FROM users WHERE token=$1"

	mRows, err := pg.db.Query(context.Background(), sql, token)
	if err != nil {
		log.Println(err)
	}

	for mRows.Next() {
		err := mRows.Scan(&u.Id, &u.Username, &u.Password, &u.Email, &u.Timezone, &u.Token)
		if err != nil {
			log.Println(err)
		}
	}

	return u, err
}

func (pg *Postgres) UpdateUserTimezone(token, timezone string) error {

	sql := "UPDATE users SET timezone=$1 WHERE token=$2"

	_, err := pg.db.Query(context.Background(), sql, timezone, token)
	if err != nil {
		log.Println(err)
	}

	return err
}

func (pg *Postgres) UpdateUserToken(token string, id int) error {

	sql := "UPDATE users SET token=$1 WHERE id=$2"

	_, err := pg.db.Query(context.Background(), sql, token, id)
	if err != nil {
		log.Println(err)
	}

	return err
}
