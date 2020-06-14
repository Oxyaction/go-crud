package user

import (
	"time"

	"github.com/google/uuid"
)

type UserCreate struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	Id        uuid.UUID
	Name      string
	Email     string
	CreatedAt time.Time
}
