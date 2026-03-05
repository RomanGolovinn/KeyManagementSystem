package postgres

import (
	"context"
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
		return nil, fmt.Errorf("Feiled to open db: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("Feiled to ping db: %w", err)
	}

	fmt.Println("DB connected")

	return &Postgres{
		DB: db,
	}, nil
}

func (p *Postgres) GetPermissions(ctx context.Context, client, server string) (int, error) {
	q := "select permision from permissions where client_id = $1 and server_id = $2"

	var permission int
	err := p.DB.QueryRowContext(ctx, q, client, server).Scan(&permission)
	if err != nil {
		return permission, fmt.Errorf("Failed to get permissions from db: %w", err)
	}
	return permission, nil

}
