package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/nkchakradhari780/practice9/internal/modules"
	"github.com/nkchakradhari780/practice9/internal/services"
	"github.com/nkchakradhari780/practice9/internal/utils/custerrors"
	"github.com/nkchakradhari780/practice9/internal/utils/response"
)

func CreateUserHandler(us services.UsersService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user modules.CreateUser

		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			slog.Error("error creating user", "error", err.Error())
			if err == io.EOF {
				response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("empty body: %v", err)))
				return
			}

			var typeError *json.UnmarshalTypeError
			if errors.As(err, &typeError) {
				response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("field %s must be of type %s", typeError.Field, typeError.Type)))
				return
			}

			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		lastId, err := us.CreateUserService(&user)
		if err != nil {
			if custerrors.IsValidationError(err) {
				validateErrors := custerrors.GetValidationErrors(err)
				response.WriteJson(w, http.StatusBadRequest, response.ValidationError(validateErrors))
				return
			}

			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return 
		}

		response.WriteJson(w, http.StatusCreated, response.GeneralSuccess(map[string]uuid.UUID{"ID": lastId}))
	}
}

func GetUserByIdHandler(us services.UsersService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")

		id, err := uuid.Parse(idStr)
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		user, err := us.GetUserByIdService(id)
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		response.WriteJson(w, http.StatusOK, response.GeneralSuccess(map[string]any{"user": user}))
	}
}