package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/nkchakradhari780/practice9/internal/api/handlers"
	"github.com/nkchakradhari780/practice9/internal/api/middlewares"
	"github.com/nkchakradhari780/practice9/internal/services"
)

func RegisterUsersRoute(mux *chi.Mux, userService services.UsersService) {
	mux.Route("/users", func(r chi.Router){
		r.Use(middlewares.AuthMiddleware) // to protect entire /users route
		r.Post("/", handlers.CreateUserHandler(userService))
		// r.With(middlewares.AuthMiddleware).Post("/", handlers.CreateUserHandler(userService))  // to protect only a single route

		//**** To protect group of router outside it will not be protected ****//
		// r.Group(func(r chi.Router){
		// 	r.Use(middlewares.AuthMiddleware)
		// 	// endpoints
		// })
	})
}