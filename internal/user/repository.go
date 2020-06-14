package user

import (
	"context"
	"fmt"

	"github.com/Oxyaction/go-crud/internal/platform/db"
	"github.com/google/uuid"
	pgx "github.com/jackc/pgx/v4"
)

type UserRepository struct {
	db *db.DB
}

func NewRepository(db *db.DB) *UserRepository {
	return &UserRepository{db}
}

func (r UserRepository) Add(ctx context.Context, user UserCreate) (*User, error) {
	newUser := User{
		Email: user.Email,
		Name:  user.Name,
	}

	err := r.db.InTx(ctx, pgx.ReadCommitted, func(tx pgx.Tx) error {
		row := tx.QueryRow(ctx, `
		INSERT INTO
			 "user"
			(name, email, password)
		VALUES
			($1, $2, $3)
		RETURNING id, createdAt
		`, user.Name, user.Email, user.Password)

		var rawUuid string

		if err := row.Scan(&rawUuid, &newUser.CreatedAt); err != nil {
			return fmt.Errorf("fetching inserted user: %w", err)
		}

		id, err := uuid.Parse(rawUuid)
		if err != nil {
			return fmt.Errorf("parsing id for inserted user: %w", err)
		}

		newUser.Id = id

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &newUser, nil
}
