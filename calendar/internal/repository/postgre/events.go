package postgre

import (
	"context"
	"http/internal/models"
	"log"
)

func (pg *Postgres) GetEvents() ([]models.Event, error) {

	var e models.Event
	sql := "SELECT * FROM events"

	pgConn := NewRepository()
	mRows, err := pgConn.Query(context.Background(), sql)
	if err != nil {
		log.Println(err)
	}

	var mSlice = make([]models.Event, 0)
	for mRows.Next() {
		err := mRows.Scan(&e.Id, &e.Title, &e.Description, &e.TimestampFrom, &e.TimestampTo)
		if err != nil {
			log.Println(err)
		}
		mSlice = append(mSlice, e)
	}

	return mSlice, err
}

func (pg *Postgres) GetEventById(id string) (models.Event, error) {

	var e models.Event
	sql := "SELECT title, description, timestamp_from, timestamp_to FROM events WHERE id=$1"

	pgConn := NewRepository()
	mRows, err := pgConn.Query(context.Background(), sql, id)
	if err != nil {
		log.Println(err)
	}

	for mRows.Next() {
		err := mRows.Scan(&e.Title, &e.Description, &e.TimestampFrom, &e.TimestampTo)
		if err != nil {
			log.Println(err)
		}
	}

	return e, err
}

func (pg *Postgres) CreateEvent(event models.Event) (int, error) {

	sql := "INSERT INTO events(user_id, title, description, timestamp_from, timestamp_to ) " +
		"VALUES ($1, $2, $3, $4, $5) RETURNING id"

	pgConn := NewRepository()
	err := pgConn.QueryRow(context.Background(), sql,
		event.Uid,
		event.Title,
		event.Description,
		event.TimestampFrom,
		event.TimestampTo,
	).Scan(&event.Id)

	if err != nil {
		log.Println(err)
	}

	return event.Id, err
}
