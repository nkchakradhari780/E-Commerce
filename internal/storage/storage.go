package storage

import (
	"github.com/nkchakradhari780/practice9/internal/modules"
)

type Storage interface {
	CreateUserDB(user *modules.CreateUser) error 
}