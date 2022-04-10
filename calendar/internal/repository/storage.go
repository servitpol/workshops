package repository

type Storage interface {
	GetUser(string) (string, error)
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
