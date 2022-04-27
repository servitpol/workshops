package postgre

import (
	"calendar/internal/models"
	"context"
	"log"
)

func (r *Repository) GetEvents() ([]models.Event, error) {

	var e models.Event
	sql := "SELECT * FROM events"

	mRows, err := r.Pool.Query(context.Background(), sql)
	if err != nil {
		log.Println(err)
	}

	var mSlice = make([]models.Event, 0)
	for mRows.Next() {
		err := mRows.Scan(&e.Id, &e.Uid, &e.Title, &e.Description, &e.TimestampFrom, &e.TimestampTo)
		if err != nil {
			log.Println(err)
		}
		mSlice = append(mSlice, e)
	}

	return mSlice, err
}

func (r *Repository) GetEventById(id string) (models.Event, error) {

	var e models.Event
	sql := "SELECT title, description, timestamp_from, timestamp_to FROM events WHERE id=$1"

	mRows, err := r.Pool.Query(context.Background(), sql, id)
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

func (r *Repository) CreateEvent(event models.Event) (int, error) {

	sql := "INSERT INTO events(user_id, title, description, timestamp_from, timestamp_to ) " +
		"VALUES ($1, $2, $3, $4, $5) RETURNING id"

	err := r.Pool.QueryRow(context.Background(), sql,
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

func (r *Repository) UpdateEvent(event models.Event, id int) error {

	sql := "UPDATE events SET title=$1, description=$2, timestamp_from=$3, timestamp_to=$4 WHERE id=$5 AND user_id=$6"

	_, err := r.Pool.Query(context.Background(), sql,
		event.Title,
		event.Description,
		event.TimestampFrom,
		event.TimestampTo,
		id,
		event.Uid,
	)

	if err != nil {
		log.Println(err)
	}

	return err
}
