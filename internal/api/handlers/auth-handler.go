package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"time"

	"github.com/nkchakradhari780/practice9/internal/modules"
	"github.com/nkchakradhari780/practice9/internal/services"
	"github.com/nkchakradhari780/practice9/internal/utils/custerrors"
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

		slog.Info("logging body", slog.Any("loginReq", loginReq))

		user, err := authService.UserLoginService(&loginReq)
		if err != nil {
			if custerrors.IsValidationError(err) {
				slog.Info("validation error")
				validateErrors := custerrors.GetValidationErrors(err)
				response.WriteJson(w, http.StatusBadRequest, response.ValidationError(validateErrors))
				return
			}
			slog.Info("user login error")
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		fmt.Println("login Response: ", user)		

		accessToken, err := authService.GenerateAccessToken(user.Id, 30 * time.Minute)
		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}

		refreshToken, err := authService.GenerateRefreshToken(user.Id, 30 * 24 * time.Hour)
		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}

		cookie := &http.Cookie{
			Name: "refresh_token", 
			Value: refreshToken,
			Path: "/",
			HttpOnly: true,
			Secure: false,
			SameSite: http.SameSiteLaxMode,
			Expires: time.Now().Add(30 * 24 * time.Hour),
		}
		http.SetCookie(w, cookie)

		response.WriteJson(w, http.StatusOK, response.GeneralSuccess(modules.LoginResponse{
			Id: user.Id,
			Name: user.Name,
			Email: user.Email,
			Role: user.Role,
			Address: user.Address,
			AccessToken: accessToken,
		}))
	}
}

func RefreshHandler(authService services.AuthService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("refresh_token")
		if err != nil {
			if errors.Is(err, http.ErrNoCookie) {
				response.WriteJson(w, http.StatusUnauthorized, response.GeneralError(fmt.Errorf("refresh token not found")))
				return
			}
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("failed to read cookie")))
			return
		}

		refreshToken := cookie.Value

		_, err = authService.RefreshUserService(refreshToken)
		if err != nil {
			slog.Error("error","", err.Error())
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

	}
}