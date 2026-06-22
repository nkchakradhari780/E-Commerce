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

func IsValidationError(err error) bool {
	_, ok := err.(*ValidationError)
	return ok
}

func GetValidationErrors(err error) validator.ValidationErrors {
	if ve, ok := err.(*ValidationError); ok {
		return ve.Errors
	}
	return nil
}