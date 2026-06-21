package modules

import "github.com/google/uuid"

type CreateUser struct {
	Id       uuid.UUID `json:"id,omitempty"`
	Name     string    `json:"name,omitempty" validate:"required"`
	Email    string    `json:"email,omitempty" validate:"required,email"`
	Password string    `json:"password,omitempty" vlidate:"required"`
	Role     string    `json:"role,omitempty" validate:"required"`
	Address  string    `json:"address,omitempty" validate:"required"`
}
