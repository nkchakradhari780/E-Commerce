package middlewares

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"strings"

	"github.com/nkchakradhari780/practice9/internal/utils/custjwt"
)

type contextKey string

const (
	TokenHeader            = "Authorization"
	TokenScheme            = "Bearer"
	userIdKey   contextKey = "user_id"
)

type ErrsResp struct {
	Error string
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Auth Middleware called")

		var token string

		authHeader := r.Header.Get(TokenHeader)
		if authHeader != "" {
			parts := strings.SplitN(authHeader, " ", 2)
			if len(parts) == 2 && parts[0] == TokenScheme {
				token = parts[1]
			}
		}

		if token == "" {
			slog.Error("token missing")
			w.WriteHeader(http.StatusUnauthorized)
			_ = json.NewEncoder(w).Encode(ErrsResp{Error: "missing token"})
			return
		}

		userId, err := custjwt.ValidateToken(token)
		if err != nil {
			slog.Error("invalid token", "error", err.Error())
			w.WriteHeader(http.StatusUnauthorized)	
			_ = json.NewEncoder(w).Encode(ErrsResp{Error: "invalid token"})
			return
		}

		ctx := context.WithValue(r.Context(), userIdKey, userId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
