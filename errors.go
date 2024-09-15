package main

import (
	"movies_api/internal/validation"
	"net/http"
)

func (app *Application) ErrorResponse(w http.ResponseWriter, r *http.Request, status int, message interface{}) {
	newErr := app.WriteJson(w, r, envelope{"errors": message}, nil, status)
	if newErr != nil {
		w.WriteHeader(500)
	}
}

func (app *Application) ServerErrorResponse(w http.ResponseWriter, r *http.Request) {
	message := "an error occurred on the the server"
	app.ErrorResponse(w, r, 500, message)
}

func (app *Application) NotFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "Page Not Found"
	app.ErrorResponse(w, r, 404, message)
}

func (app *Application) BadRequestResponse(w http.ResponseWriter, r *http.Request) {
	message := "The request is not written correctly"
	app.ErrorResponse(w, r, 400, message)
}

func (app *Application) MethodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	message := "The http method used is not allowed"
	app.ErrorResponse(w, r, 405, message)
}

func (app *Application) ValidationErrorResponse(w http.ResponseWriter, r *http.Request, v *validation.Validator) {
	app.ErrorResponse(w, r, http.StatusUnprocessableEntity, v.Errors)
}

func (app *Application) InvalidCredentialsResponse(w http.ResponseWriter, r *http.Request) {
	message := "invalid authentication credentials"
	app.ErrorResponse(w, r, http.StatusUnauthorized, message)
}

func (app *Application) InvalidAuthenticationTokenResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("WWW-Authenticate", "Bearer")
	message := "invalid or missing authentication token"
	app.ErrorResponse(w, r, http.StatusUnauthorized, message)
}

func (app *Application) AuthenticationRequiredResponse(w http.ResponseWriter, r *http.Request) {
	message := "you must be authenticated to access this resource"
	app.ErrorResponse(w, r, http.StatusUnauthorized, message)
}

func (app *Application) InactiveAccountResponse(w http.ResponseWriter, r *http.Request) {
	message := "your user account must be activated to access this resource"
	app.ErrorResponse(w, r, http.StatusForbidden, message)
}

func (app *Application) NotPermittedResponse(w http.ResponseWriter, r *http.Request) {
	message := "your user account doesn't have the necessary permissions to access this resource"
	app.ErrorResponse(w, r, http.StatusForbidden, message)
}
