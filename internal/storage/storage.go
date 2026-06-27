package storage

import (
	"time"

	"github.com/google/uuid"
	"github.com/nkchakradhari780/practice9/internal/modules"
)

type Storage interface {
	// Auth DB
	CreateRefreshToken(userId uuid.UUID, tokenHash string, expiresAt time.Time) error

	// User DB
	CreateUserDB(user *modules.CreateUser) error 
	GetUserByEmail(email string) (*modules.User, error)

}