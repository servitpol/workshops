package postgre

import (
	"context"
	"fmt"
	"github.com/jackc/pgconn"
	"http/internal/config"
	"log"
)

type Postgres struct {
	//
}

func NewRepository() *pgconn.PgConn {

	cfg := config.GetConfig()
	var dbUrl = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", cfg.DB.User, cfg.DB.Pass, cfg.DB.Host, cfg.DB.Port, cfg.DB.DbName)

	pgConnect, err := pgconn.Connect(context.Background(), dbUrl)
	if err != nil {
		log.Fatalln("pgСonn failed to connect:", err)
	}

	return pgConnect
}
