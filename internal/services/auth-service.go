package services

import (
	"errors"

	"github.com/go-playground/validator"
	"github.com/nkchakradhari780/practice9/internal/modules"
	"github.com/nkchakradhari780/practice9/internal/storage"
	"github.com/nkchakradhari780/practice9/internal/utils/custerrors"
	"github.com/nkchakradhari780/practice9/internal/utils/hashpass"
)

type AuthService interface {
	UserLoginService(logReq *modules.LoginUser) (*modules.User, error)
}

type authService struct {
	storage storage.Storage
}

func NewAuthService(storage storage.Storage) AuthService {
	return &authService{storage: storage}
}

func (ls *authService) UserLoginService(logReq *modules.LoginUser) (*modules.User, error) {
	if err := validator.New().Struct(&logReq); err != nil {
		var validatorError validator.ValidationErrors
		if errors.As(err, &validatorError) {
			return nil, custerrors.NewValidationError(validatorError)
		}
		return nil, err
	}

	user, err := ls.storage.GetUserByEmail(logReq.Email)
	if err != nil {
		return nil, err
	}

	if err := hashpass.CompareHash(logReq.Password, user.Password); err != nil {
		return nil, err
	}

	return user, nil
	
}