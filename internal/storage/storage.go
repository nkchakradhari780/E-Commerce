package storage

import "github.com/google/uuid"

type Storage interface {
	CreateUserDB(id uuid.UUID, name, email, password, role, address string) error
}