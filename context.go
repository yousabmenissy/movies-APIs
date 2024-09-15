package main

import (
	"context"
	"movies_api/data/users"
	"net/http"
)

type ContextKey string

const userContextKey = ContextKey("user")

func (app *Application) ContextSetUser(r *http.Request, user *users.User) *http.Request {
	return r.WithContext(context.WithValue(r.Context(), userContextKey, user))
}

func (app *Application) contextGetUser(r *http.Request) *users.User {
	user, ok := r.Context().Value(userContextKey).(*users.User)
	if !ok {
		panic("missing user value in request context")
	}
	return user
}
