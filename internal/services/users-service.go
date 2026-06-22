package services

import (
	"github.com/nkchakradhari780/practice9/internal/modules"
	"github.com/nkchakradhari780/practice9/internal/storage"
)

type UsersService interface {
	CreateUserService(user *modules.CreateUser) error
}

type usersService struct {
	storage storage.Storage
}

func NewUserService(storage storage.Storage) UsersService {
	return &usersService{storage: storage}
}

func (us *usersService) CreateUserService(user *modules.CreateUser) error {
	
	return nil
}