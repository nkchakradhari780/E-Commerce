package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/nkchakradhari780/practice9/internal/api/handlers"
	"github.com/nkchakradhari780/practice9/internal/services"
)

func RegisterUsersRoute(mux *chi.Mux, userService services.UsersService) {
	mux.Route("/users", func(r chi.Router){
		r.Post("/", handlers.CreateUserHandler(userService))
	})
}