package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/nkchakradhari780/practice9/internal/modules"
	"github.com/nkchakradhari780/practice9/internal/services"
	"github.com/nkchakradhari780/practice9/internal/utils/custerrors"
	"github.com/nkchakradhari780/practice9/internal/utils/custjwt"
	"github.com/nkchakradhari780/practice9/internal/utils/response"
)

func LoingUserHandler(authService services.AuthService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var loginReq modules.LoginUser

		if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
			if err == io.EOF{
				response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("empty body: %v", err)))
				return
			}

			var typeError *json.UnmarshalTypeError
			if errors.As(err, &typeError) {
				response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("field %s should be of type %s", typeError.Field, typeError.Type)))
				return
			}
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		user, err := authService.UserLoginService(&loginReq)
		if err != nil {
			if custerrors.IsValidationError(err) {
				validateErrors := custerrors.GetValidationErrors(err)
				response.WriteJson(w, http.StatusBadRequest, response.ValidationError(validateErrors))
				return
			}
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		fmt.Println("login Response: ", user)		

		token, err := custjwt.GenerateToken(user.Id, 24 * time.Hour)
		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}

		cookie := &http.Cookie{
			Name: "auth_token", 
			Value: token, 
			Path: "/",
			HttpOnly: true,
			Secure: false, 
			SameSite: http.SameSiteLaxMode,
			Expires: time.Now().Add(24 * time.Hour), 
		}

		http.SetCookie(w, cookie)

		response.WriteJson(w, http.StatusOK, response.GeneralSuccess(modules.LoginResponse{
			Id: user.Id,
			Name: user.Name,
			Email: user.Email,
			Role: user.Role,
			Address: user.Address,
			Token: token,
		}))
	}
}