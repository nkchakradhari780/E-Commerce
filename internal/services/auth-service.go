package services

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"log/slog"
	"time"

	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"github.com/nkchakradhari780/practice9/internal/modules"
	"github.com/nkchakradhari780/practice9/internal/storage"
	"github.com/nkchakradhari780/practice9/internal/utils/custerrors"
	"github.com/nkchakradhari780/practice9/internal/utils/custjwt"
	"github.com/nkchakradhari780/practice9/internal/utils/hashpass"
)

type AuthService interface {
	UserLoginService(logReq *modules.LoginUser) (*modules.User, error)
	GenerateRefreshToken(userId uuid.UUID, ttl time.Duration) (string, error)
	GenerateAccessToken(userId uuid.UUID, ttl time.Duration) (string, error)
}

type authService struct {
	storage storage.Storage
}

func NewAuthService(storage storage.Storage) AuthService {
	return &authService{storage: storage}
}

func (ls *authService) UserLoginService(logReq *modules.LoginUser) (*modules.User, error) {
	if err := validator.New().Struct(logReq); err != nil {
		slog.Error("validation error")
		var validatorError validator.ValidationErrors
		if errors.As(err, &validatorError) {
			return nil, custerrors.NewValidationError(validatorError)
		}
		return nil, err
	}

	user, err := ls.storage.GetUserByEmailDB(logReq.Email)
	if err != nil {
		slog.Error("email scan error", "error", err.Error())
		return nil, err
	}

	if err := hashpass.CompareHash(logReq.Password, user.Password); err != nil {
		return nil, err
	}

	return user, nil	
}

func (ls *authService) GenerateRefreshToken(userId uuid.UUID, ttl time.Duration) (string, error) {
	refreshToken, err := custjwt.GenerateRefreshToken(userId, ttl)
	if err != nil {
		return "", err
	}

	tokenHash:= sha256.Sum256([]byte(refreshToken))
	tokenHashStr := hex.EncodeToString(tokenHash[:])
	expiresAt := time.Now().Add(ttl).UTC().Local()
	tokenId := uuid.New()

	if err := ls.storage.CreateRefreshTokenDB(tokenId, userId, tokenHashStr, expiresAt); err != nil {
		return "", err
	}

	return refreshToken, nil
}

func (ls *authService) GenerateAccessToken(userId uuid.UUID, ttl time.Duration) (string, error) {
	accessToken, err := custjwt.GenerateAccessToken(userId, ttl)
	if err != nil {
		return "", err
	}
	return accessToken, nil 
}