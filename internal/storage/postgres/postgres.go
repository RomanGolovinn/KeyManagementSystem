package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres struct {
	DB *pgxpool.Pool
}

func NewPostgres(ctx context.Context, dsn string) (*Postgres, error) {
	db, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to postgres: %w", err)
	}

	err = db.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("Failed to ping postgres: %w", err)
	}

	return &Postgres{
		DB: db,
	}, nil
}

func (p *Postgres) GetPermissions(ctx context.Context, client, server string) (PermissionType, error) {
	q := "select permission from permissions where client_id = $1 and server_id = $2"

	var permission PermissionType
	err := p.DB.QueryRow(ctx, q, client, server).Scan(&permission)
	if err != nil {
		return permission, fmt.Errorf("Failed to get permissions from db: %w", err)
	}
	return permission, nil

}

func (p *Postgres) CreateClient(ctx context.Context, client Client) (string, error) {
	q := `insert into clients (name, tls_pem, valid_to) values ($1, $2, $3)
		returning id`

	var id string

	err := p.DB.QueryRow(ctx, q, client.Name, client.TlsPem,
		client.ValidTo).Scan(&id)

	if err != nil {
		return "", fmt.Errorf("Failed to create client in db: %w", err)
	}

	return id, nil
}

func (p *Postgres) CreateServer(ctx context.Context, server Server) (string, error) {
	q := `insert into servers (name, address) values ($1, $2)
		returning id`

	var id string

	err := p.DB.QueryRow(ctx, q, server.Name, server.Address).Scan(&id)
	if err != nil {
		return "", fmt.Errorf("Failed to create server in db: %w", err)
	}

	return id, nil
}

func (p *Postgres) CreatePermission(ctx context.Context, permission Permission) error {
	q := `insert into permissions (client_id, server_id, permission) values ($1, $2, $3)`

	_, err := p.DB.Exec(ctx, q, permission.ClientId, permission.ServerId,
		permission.Permission)
	if err != nil {
		return fmt.Errorf("Failed to create permission in db: %w", err)
	}
	return nil
}

func (p *Postgres) CreateIssuedCredential(ctx context.Context, credential IssuedCredential) error {
	q := `insert into issued_credentials (client_id, server_id, permission, temp_username, 
		temp_password_hash, created_at, valid_to)
		values ($1, $2, $3, $4, $5, $6, $7) returning id`

	var id string

	err := p.DB.QueryRow(ctx, q, credential.ClientId, credential.ServerId, credential.Permission,
		credential.TempUsername, credential.TempPasswordHash, credential.CreatedAt, credential.ValidTo).Scan(&id)
	if err != nil {
		return fmt.Errorf("Failed to create issued credential in db: %w", err)
	}
	return nil
}

func (p *Postgres) RevokeIssuedCredential(ctx context.Context, credentialId string, revokedAt time.Time) error {
	q := `update issued_credentials set revoked_at = $1 where id = $2`

	_, err := p.DB.Exec(ctx, q, revokedAt, credentialId)
	if err != nil {
		return fmt.Errorf("Failed to revoke issued credential in db: %w", err)
	}
	return nil
}

func (p *Postgres) CreateAuditLog(ctx context.Context, log AuditLog) error {
	q := `insert into audit_log (timestamp, client_id, server_id, event_type, ip_address, details)
		values ($1, $2, $3, $4, $5, $6)`

	// If we register an unknown client without a UUID
	// the database cannot write an empty string to a UUID-type field
	// but it can write a NULL (if the pointer is nil).
	var clientID *string
	if log.ClientId != "" {
		clientID = &log.ClientId
	}

	var serverID *string
	if log.ServerId != "" {
		serverID = &log.ServerId
	}

	_, err := p.DB.Exec(ctx, q, log.Timestamp, clientID, serverID,
		log.EventType, log.IpAddress, log.Details)
	if err != nil {
		return fmt.Errorf("Failed to create audit log in db: %w", err)
	}
	return nil
}
