package storage

import (
	"time"

	"github.com/google/uuid"
	"github.com/nkchakradhari780/practice9/internal/modules"
)

type Storage interface {
	// Auth DB
	CreateRefreshTokenDB(tokenId, userId uuid.UUID, tokenHash string, expiresAt time.Time) error

	// User DB
	CreateUserDB(user *modules.CreateUser) error 
	GetUserByEmailDB(email string) (*modules.User, error)
	GetUserByIdDB(userId uuid.UUID) (*modules.GetUser, error)
}