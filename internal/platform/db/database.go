package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type DbConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

type DB struct {
	Pool *pgxpool.Pool
}

type Defer func()

func NewDB(ctx context.Context, config *DbConfig) (*DB, Defer, error) {
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
	)
	pool, err := pgxpool.Connect(ctx, connectionString)
	if err != nil {
		return nil, nil, fmt.Errorf("creating connection pool: %v", err)
	}

	db := &DB{
		Pool: pool,
	}

	return db, func() { db.Close() }, nil
}

func (db *DB) InTx(ctx context.Context, isoLevel pgx.TxIsoLevel, f func(tx pgx.Tx) error) error {
	conn, err := db.Pool.Acquire(ctx)
	if err != nil {
		return fmt.Errorf("acquiring connection: %v", err)
	}
	defer conn.Release()

	tx, err := conn.BeginTx(ctx, pgx.TxOptions{IsoLevel: isoLevel})
	if err != nil {
		return fmt.Errorf("starting transaction: %v", err)
	}

	if err := f(tx); err != nil {
		if err1 := tx.Rollback(ctx); err1 != nil {
			return fmt.Errorf("rolling back transaction: %v (original error: %v)", err1, err)
		}
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("committing transaction: %v", err)
	}
	return nil
}

func (db *DB) Close() {
	db.Pool.Close()
}
