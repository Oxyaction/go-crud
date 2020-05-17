package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
)

type DbConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

func Open(ctx context.Context, config *DbConfig) (*pgx.Conn, error) {
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
	)
	return pgx.Connect(ctx, connectionString)
}
