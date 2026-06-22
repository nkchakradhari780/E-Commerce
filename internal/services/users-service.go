package services

import (
	"errors"

	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"github.com/nkchakradhari780/practice9/internal/modules"
	"github.com/nkchakradhari780/practice9/internal/storage"
	"github.com/nkchakradhari780/practice9/internal/utils/custerrors"
	"github.com/nkchakradhari780/practice9/internal/utils/hashpass"
)

type UsersService interface {
	CreateUserService(user *modules.CreateUser) (uuid.UUID, error)
}

type usersService struct {
	storage storage.Storage
}

func NewUserService(storage storage.Storage) UsersService {
	return &usersService{storage: storage}
}

func (us *usersService) CreateUserService(user *modules.CreateUser) (uuid.UUID, error) {
	if err := validator.New().Struct(user); err != nil {
		var validationError validator.ValidationErrors
		if errors.As(err, &validationError) {
			return uuid.Nil, custerrors.NewValidationError(validationError)
		}
		return uuid.Nil, err
	}

	uid := uuid.New()

	hashed, err := hashpass.GenerateHash(user.Password)
	if err != nil {
		return uuid.Nil, err
	}

	user.Id = uid
	user.Password = hashed

	if err := us.storage.CreateUserDB(user)
	err != nil {
		return uuid.Nil, err
	}
	 
	return uid, nil
}