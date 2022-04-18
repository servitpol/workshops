package repository

import "http/internal/models"

type Storage interface {
	GetUserByUsername(string) (models.User, error)
	GetUserByToken(string) (models.User, error)
	UpdateUserTimezone(token, timezone string) error
	UpdateUserToken(token string, id int) error
	GetEvents() ([]models.Event, error)
	GetEventById(string) (models.Event, error)
	CreateEvent(event models.Event) (int, error)
}

type Db struct {
	Db Storage
}

func NewStorage(s Storage) *Db {
	return &Db{
		Db: s,
	}
}
