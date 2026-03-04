package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type Postgres struct {
	DB *sql.DB
}

func NewPostgres(dsn string) (*Postgres, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("Feiled to open db: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("Feiled to ping db: %v", err)
	}

	fmt.Println("DB connected")

	return &Postgres{
		DB: db,
	}, nil
}
