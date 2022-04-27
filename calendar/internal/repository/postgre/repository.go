package postgre

import (
	"calendar/internal/config"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

type Repository struct {
	Pool *pgxpool.Pool
}

func New() *pgxpool.Pool {

	cfg := config.GetConfig()
	var dbUrl = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", cfg.DB.User, cfg.DB.Pass, cfg.DB.Host, cfg.DB.Port, cfg.DB.DbName)

	pgConnect, err := pgxpool.Connect(context.Background(), dbUrl)
	if err != nil {
		log.Fatalln("pgСonn failed to connect:", err)
	}

	return pgConnect
}
