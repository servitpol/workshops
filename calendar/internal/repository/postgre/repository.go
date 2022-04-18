package postgre

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"http/internal/config"
	"log"
)

type Postgres struct {
	db *pgxpool.Pool
}

func New() *Postgres {

	cfg := config.GetConfig()
	var dbUrl = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", cfg.DB.User, cfg.DB.Pass, cfg.DB.Host, cfg.DB.Port, cfg.DB.DbName)

	pgConnect, err := pgxpool.Connect(context.Background(), dbUrl)
	if err != nil {
		log.Fatalln("pgСonn failed to connect:", err)
	}

	return &Postgres{db: pgConnect}
}
