package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator"
)

type Response struct {
	Status string
	Error string 
	Data any
}

const (
	StatusError = "Error"
	StatusOk = "Ok"
)

func WriteJson(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "Application/Json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

func GeneralError(err error) *Response {
	return &Response{
		Status: StatusError,
		Error: err.Error(),
		Data: nil,
	}
}

func GeneralSuccess(data any) *Response {
	return &Response{
		Status: StatusOk, 
		Data: data,
	}
}

func ValidationError(errs validator.ValidationErrors) *Response {
	var errsMap []string

	for _, err := range errs {
		switch err.ActualTag() {
		case "required":
			errsMap = append(errsMap, fmt.Sprintf("field %s is a required field", err.Field()))
		default:
			errsMap = append(errsMap, fmt.Sprintf("field %s is invalid field", err.Field()))
		}
	}


	return &Response{
		Status: StatusError,
		Error: strings.Join(errsMap, ","),
		Data: nil,
	}
}