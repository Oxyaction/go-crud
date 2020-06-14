package user

import (
	"context"

	"github.com/Oxyaction/go-crud/internal/platform/db"
)

type UserManager struct {
	repository *UserRepository
}

func NewManager(db *db.DB) *UserManager {
	return &UserManager{
		repository: NewRepository(db),
	}
}

func (m UserManager) Register(ctx context.Context, userToCreate UserCreate) (*User, error) {
	return m.repository.Add(ctx, userToCreate)
}
