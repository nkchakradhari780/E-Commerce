package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/nkchakradhari780/practice9/internal/api/handlers"
	"github.com/nkchakradhari780/practice9/internal/services"
)

func RegisterAuthRoute(mux *chi.Mux, authService services.AuthService) {
	mux.Route("/auth", func(r chi.Router){
		r.Post("/login", handlers.LoingUserHandler(authService))
		r.Post("/refresh", handlers.RefreshHandler(authService))
	})
}