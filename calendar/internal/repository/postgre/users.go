package postgre

import (
	"context"
	"http/internal/models"
	"log"
)

func (pg *Postgres) GetUserByUsername(username string) (models.User, error) {

	var u models.User
	sql := "SELECT * FROM users WHERE username=$1"

	pgConn := NewRepository()
	mRows, err := pgConn.Query(context.Background(), sql, username)
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

	pgConn := NewRepository()
	mRows, err := pgConn.Query(context.Background(), sql, token)
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

func (pg *Postgres) UpdateUserToken(token string) (models.User, error) {

	var u models.User
	sql := "SELECT * FROM users WHERE token=$1"

	pgConn := NewRepository()
	mRows, err := pgConn.Query(context.Background(), sql, token)
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
