package repository

import "http/internal/handlers"

type Storage interface {
	Login(user *handlers.User) error
	GetEvents() error
}

type db struct {
	Db Storage
}

func NewStorage(st Storage) *db {
	return &db{
		Db: st,
	}
}
