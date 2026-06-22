package custerrors

import "github.com/go-playground/validator"

type ValidationError struct {
	Errors validator.ValidationErrors
}

func (err *ValidationError) Error() string {
	return "validation error"
}

func NewValidationError(errs validator.ValidationErrors) *ValidationError {
	return &ValidationError{Errors: errs}
}

