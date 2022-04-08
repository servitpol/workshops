package postgre

import (
	"context"
	"fmt"
	"log"
)

func (pg *Postgres) GetEvents() error {

	sql := "SELECT * FROM events"
	pgConn := NewRepository()

	var paramValues [][]byte
	result := pgConn.ExecParams(context.Background(), sql, paramValues, nil, nil, nil)
	for result.NextRow() {
		fmt.Println("Events:", string(result.Values()[0]))
	}

	_, err := result.Close()
	if err != nil {
		log.Fatalln("failed reading result:", err)
	}

	return err
}
