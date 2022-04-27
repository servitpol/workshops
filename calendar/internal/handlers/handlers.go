package handlers

import (
	"calendar/internal/repository"
)

type Handler struct {
	Storage repository.Db
}
