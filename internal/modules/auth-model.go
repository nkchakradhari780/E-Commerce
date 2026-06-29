package modules

import (
	"time"

	"github.com/google/uuid"
)

type LoginUser struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Id           uuid.UUID `json:"id,omitempty"`
	Name         string    `json:"name,omitempty"`
	Email        string    `json:"email,omitempty"`
	Role         string    `json:"role,omitempty"`
	Address      string    `json:"address,omitempty"`
	AccessToken  string    `json:"acccess_token,omitempty"`
	RefreshToken string    `json:"refresh_token,omitempty"`
}

type RefreshToken struct {
	Id        uuid.UUID `json:"id,omitempty"`
	Revoked   bool      `json:"revoked,omitempty"`
	ExpiresAt time.Time `json:"expires_at,omitempty"`
}

type UserId struct {
	Id uuid.UUID `json:"user_id" validate:"required"`
}