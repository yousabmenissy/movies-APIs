package main

import (
	"errors"
	"movies_api/data/tokens"
	"movies_api/data/users"
	"movies_api/internal"
	"movies_api/internal/validation"
	"net/http"
	"strings"
)

func (app *Application) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Vary", "Authorization")

		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			r = app.ContextSetUser(r, users.AnnonymousUser)
			next.ServeHTTP(w, r)
			return
		}

		headerParts := strings.Split(authHeader, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			app.InvalidAuthenticationTokenResponse(w, r)
			return
		}

		token := headerParts[1]

		v := validation.New()

		if tokens.ValidateTokenPlaintext(&v, token); !v.Valid() {
			app.InvalidAuthenticationTokenResponse(w, r)
			return
		}

		user, err := app.models.Users.GetForToken(token, tokens.ScopeAuthentication)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrNoRecord):
				app.InvalidAuthenticationTokenResponse(w, r)
			default:
				app.ServerErrorResponse(w, r)
			}
			return
		}

		r = app.ContextSetUser(r, user)

		next.ServeHTTP(w, r)
	})
}
func (app *Application) RequireActivatedUser(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := app.contextGetUser(r)

		if user.IsAnonymous() {
			app.AuthenticationRequiredResponse(w, r)
			return
		}

		if !user.Activated {
			app.InactiveAccountResponse(w, r)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (app *Application) RequirePermission(code string, next http.HandlerFunc) http.HandlerFunc {
	fn := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := app.contextGetUser(r)

		permissions, err := app.models.Permissions.GetAllForUser(user.Id)
		if err != nil {
			app.ServerErrorResponse(w, r)
			return
		}

		if !permissions.Include(code) {
			app.NotPermittedResponse(w, r)
			return
		}

		next.ServeHTTP(w, r)
	})
	return app.RequireActivatedUser(fn)
}
