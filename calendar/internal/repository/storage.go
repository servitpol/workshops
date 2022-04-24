package repository

import "http/internal/models"

//go:generate mockgen -source=storage.go -destination=mocks/mock.go

type Storage interface {
	GetUserByUsername(string) (models.User, error)
	GetUserByToken(string) (models.User, error)
	UpdateUserTimezone(token, timezone string) error
	UpdateUserToken(token string, id int) error
	GetEvents() ([]models.Event, error)
	GetEventById(string) (models.Event, error)
	CreateEvent(event models.Event) (int, error)
	UpdateEvent(event models.Event, id int) error
}

type Db struct {
	Db Storage
}

func NewStorage(st Storage) *Db {
	return &Db{
		Db: st,
	}
}
