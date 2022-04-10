package postgre

import (
	"context"
	"log"
)

func (pg *Postgres) GetUser(username string) (string, error) {

	sql := "SELECT password FROM users WHERE username=$1"

	paramValues := [][]byte{
		[]byte(username),
	}
	pgConn := NewRepository()

	result := pgConn.ExecParams(context.Background(), sql, paramValues, nil, nil, nil)
	var pass string
	for result.NextRow() {
		pass = string(result.Values()[0])
	}
	_, err := result.Close()
	if err != nil {
		log.Fatalln("failed reading result:", err)
	}

	return pass, err
}
