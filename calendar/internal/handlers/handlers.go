package handlers

import (
	"http/internal/repository"
)

type Handler struct {
	Storage repository.Db
}
