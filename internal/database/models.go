// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0

package database

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UserID       uuid.UUID
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Email        string
	PasswordHash string
}
