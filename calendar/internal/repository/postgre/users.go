package postgre

import (
	"context"
	"fmt"
	"http/internal/handlers"
	"log"
)

func (pg *Postgres) Login(user *handlers.User) error {

	sql := "SELECT email FROM users WHERE username=$1 AND password=$2"

	paramValues := [][]byte{
		[]byte(user.Username),
		[]byte(user.Password),
	}
	pgConn := NewRepository()

	result := pgConn.ExecParams(context.Background(), sql, paramValues, nil, nil, nil)
	for result.NextRow() {
		fmt.Println("User has email:", string(result.Values()[0]))
	}
	_, err := result.Close()
	if err != nil {
		log.Fatalln("failed reading result:", err)
	}

	return err
}
