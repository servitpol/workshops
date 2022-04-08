package handlers

import (
	storage "http/internal/repository/postgre"
)

type Handler struct {
	Storage *storage.Postgres
}
