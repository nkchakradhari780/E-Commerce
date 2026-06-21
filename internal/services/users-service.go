package services

import (
	"github.com/nkchakradhari780/practice9/internal/modules"
	"github.com/nkchakradhari780/practice9/internal/storage"
)

type UsersService interface {
}

type usersService struct {
	storage storage.Storage
}

func CreateUserService(user *modules.CreateUser) error {
	
	return nil
}