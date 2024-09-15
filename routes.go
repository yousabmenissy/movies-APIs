package main

import (
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
)

func (app *Application) routes() http.Handler {
	router := http.NewServeMux()
	router.HandleFunc("GET /", app.HealthCheckHandler)

	router.HandleFunc("GET /movies", app.RequirePermission("movies:read", app.GetMoviesHandler))
	router.HandleFunc("GET /movies/{id}", app.RequirePermission("movies:read", app.GetMovieHandler))
	router.HandleFunc("POST /movies", app.RequirePermission("movies:write", app.InsertMovieHandler))
	router.HandleFunc("PUT /movies/{id}", app.RequirePermission("movies:write", app.UpdateMovieHandler))
	router.HandleFunc("DELETE /movies/{id}", app.RequirePermission("movies:write", app.DeleteMovieHandler))

	router.HandleFunc("POST /users", app.RegisterUserHandler)
	router.HandleFunc("GET /users", app.GetUserHandler)
	router.HandleFunc("PUT /users/activated", app.ActivateUserHandler)
	router.HandleFunc("POST /tokens/authentication", app.CreateAuthTokenHanler)

	return app.Authenticate(middleware.Logger(router))
}
