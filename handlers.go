package main

import (
	"database/sql"
	"errors"
	"movies_api/data/movies"
	"movies_api/data/tokens"
	"movies_api/data/users"
	"movies_api/internal"
	"movies_api/internal/validation"
	"net/http"
	"strconv"
	"time"
)

func (app *Application) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK\n"))
}

func (app *Application) GetMoviesHandler(w http.ResponseWriter, r *http.Request) {
	var params internal.URLParams
	qs := r.URL.Query()

	app.readURLParams(qs, &params)

	movies, err := app.models.Moives.GetMovies(params)
	if err != nil {
		app.logger.LogError.Println(err)
		app.ServerErrorResponse(w, r)
		return
	}

	if len(movies) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	err = app.WriteJson(w, r, envelope{"movies": movies}, http.Header{}, 200)
	if err != nil {
		app.logger.LogError.Println(err)
		app.ServerErrorResponse(w, r)
		return
	}
}

func (app *Application) GetMovieHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	movieId, err := strconv.Atoi(id)
	if err != nil {
		app.logger.LogError.Println(err)
		app.BadRequestResponse(w, r)
		return
	}
	movie, err := app.models.Moives.GetMovie(movieId)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			app.logger.LogError.Println(err)
			app.NotFoundResponse(w, r)
		default:
			app.logger.LogError.Println(err)
			app.ServerErrorResponse(w, r)
		}
		return
	}

	if err = app.WriteJson(w, r, envelope{"movie": movie}, http.Header{}, 200); err != nil {
		app.logger.LogError.Println(err)
		app.ServerErrorResponse(w, r)
		return
	}
}

func (app *Application) InsertMovieHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title          string   `json:"title"`
		Director       string   `json:"director"`
		Producers      []string `json:"producers"`
		Prod_companies []string `json:"prod_companies"`
		Writers        []string `json:"writers"`
		Overview       string   `json:"overview"`
		Status         string   `json:"status"`
		Budget         int      `json:"budget"`
		Age_rating     string   `json:"age_rating"`
		Language       string   `json:"language"`
		Runtime        int      `json:"runtime"`
		Cast_members   []string `json:"cast_members"`
		Genres         []string `json:"genres"`
		Release_date   string   `json:"release_date"`
		Country        string   `json:"country"`
	}

	err := app.ReadJson(r, &input)
	if err != nil {
		app.logger.LogError.Println(err)
		app.BadRequestResponse(w, r)
		return
	}

	movie := movies.Movie{
		Title:          input.Title,
		Director:       input.Director,
		Producers:      input.Producers,
		Prod_companies: input.Prod_companies,
		Writers:        input.Writers,
		Overview:       input.Overview,
		Status:         input.Status,
		Budget:         input.Budget,
		Age_rating:     input.Age_rating,
		Language:       input.Language,
		Runtime:        input.Runtime,
		Cast_members:   input.Cast_members,
		Genres:         input.Genres,
		Release_date:   input.Release_date,
		Country:        input.Country,
	}

	v := validation.New()
	if movies.ValidateMovie(&v, &movie); !v.Valid() {
		app.ValidationErrorResponse(w, r, &v)
		return
	}

	err = app.models.Moives.InsertMovie(&movie)
	if err != nil {
		app.logger.LogError.Println(err)
		app.ServerErrorResponse(w, r)
	}

	w.WriteHeader(http.StatusCreated)
}

func (app *Application) UpdateMovieHandler(w http.ResponseWriter, r *http.Request) {
	idstr := r.PathValue("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		app.logger.LogError.Println(err)
		app.ServerErrorResponse(w, r)
		return
	}
	var input struct {
		Title          string   `json:"title"`
		Director       string   `json:"director"`
		Producers      []string `json:"producers"`
		Prod_companies []string `json:"prod_companies"`
		Writers        []string `json:"writers"`
		Overview       string   `json:"overview"`
		Status         string   `json:"status"`
		Budget         int      `json:"budget"`
		Age_rating     string   `json:"age_rating"`
		Language       string   `json:"language"`
		Runtime        int      `json:"runtime"`
		Cast_members   []string `json:"cast_members"`
		Genres         []string `json:"genres"`
		Release_date   string   `json:"release_date"`
		Country        string   `json:"country"`
	}

	err = app.ReadJson(r, &input)
	if err != nil {
		app.logger.LogError.Println(err)
		app.BadRequestResponse(w, r)
		return
	}

	movie := movies.Movie{
		Id:             int64(id),
		Title:          input.Title,
		Director:       input.Director,
		Producers:      input.Producers,
		Prod_companies: input.Prod_companies,
		Writers:        input.Writers,
		Overview:       input.Overview,
		Status:         input.Status,
		Budget:         input.Budget,
		Age_rating:     input.Age_rating,
		Language:       input.Language,
		Runtime:        input.Runtime,
		Cast_members:   input.Cast_members,
		Genres:         input.Genres,
		Release_date:   input.Release_date,
		Country:        input.Country,
	}

	err = app.models.Moives.UpdateMovie(movie)
	if err != nil {
		app.logger.LogError.Println(err)
		app.BadRequestResponse(w, r)
		return
	}
}

func (app *Application) DeleteMovieHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	movieId, err := strconv.Atoi(id)
	if err != nil {
		app.logger.LogError.Println(err)
		app.BadRequestResponse(w, r)
		return
	}

	err = app.models.Moives.DeleteMovie(int64(movieId))
	if err != nil {
		switch {
		case errors.Is(err, internal.ErrNoRecord):
			app.NotFoundResponse(w, r)
		default:
			app.ServerErrorResponse(w, r)
		}
		app.logger.LogError.Println(err)
		return
	}
	w.WriteHeader(204)
}

func (app *Application) RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	err := app.ReadJson(r, &input)
	if err != nil {
		app.logger.LogError.Println(err)
		app.BadRequestResponse(w, r)
		return
	}

	user := users.User{
		Name:      input.Name,
		Email:     input.Email,
		Activated: false,
	}

	err = user.Password.Set(input.Password)
	if err != nil {
		app.logger.LogError.Println(err)
		app.ServerErrorResponse(w, r)
		return
	}

	v := validation.New()
	if users.ValidateUser(&v, user); !v.Valid() {
		app.ValidationErrorResponse(w, r, &v)
		return
	}

	err = app.models.Users.InserUser(&user)
	if err != nil {
		switch {
		case errors.Is(err, internal.ErrDuplicateEmail):
			v.AddError("email", "a user with this email already exists")
			app.ValidationErrorResponse(w, r, &v)
		default:
			app.ServerErrorResponse(w, r)
		}
		app.logger.LogError.Println(err)
		return
	}

	if err := app.models.Permissions.AddForUser(user.Id, "movies:read"); err != nil {
		app.logger.LogError.Println(err)
		app.ServerErrorResponse(w, r)
		return
	}

	token, err := app.models.Tokens.New(user.Id, 3*24*time.Hour, tokens.ScopeActivation)
	if err != nil {
		app.logger.LogError.Println(err)
		app.ServerErrorResponse(w, r)
		return
	}

	app.background(func() {
		data := map[string]interface{}{
			"activationToken": token.Plaintext,
			"userID":          user.Id,
		}
		err = app.mailer.Send(user.Email, "user_welcome.tmpl", data)
		if err != nil {
			app.logger.LogError.Printf("failed to send activation email to %s: %s", user.Email, err)
		}
	})

	err = app.WriteJson(w, r, envelope{"user": user}, nil, 200)
	if err != nil {
		app.logger.LogError.Println(err)
		app.ServerErrorResponse(w, r)
	}
}

func (app *Application) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	qs := r.URL.Query()
	email := app.ReadString(qs, "email", "")

	v := validation.New()
	if users.ValidateEmail(&v, email); !v.Valid() {
		app.ValidationErrorResponse(w, r, &v)
		return
	}

	user, err := app.models.Users.GetUserByEmail(email)
	if err != nil {
		switch {
		case errors.Is(err, internal.ErrNoRecord):
			app.NotFoundResponse(w, r)
		default:
			app.ServerErrorResponse(w, r)
		}
		app.logger.LogError.Println(err)
		return
	}

	if err = app.WriteJson(w, r, envelope{"user": user}, http.Header{}, 200); err != nil {
		app.logger.LogError.Println(err)
		app.ServerErrorResponse(w, r)
		return
	}
}

func (app *Application) ActivateUserHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		TokenPlaintext string `json:"token"`
	}

	err := app.ReadJson(r, &input)
	if err != nil {
		app.logger.LogError.Println(err)
		app.BadRequestResponse(w, r)
		return
	}

	v := validation.New()
	if tokens.ValidateTokenPlaintext(&v, input.TokenPlaintext); !v.Valid() {
		app.ValidationErrorResponse(w, r, &v)
		return
	}

	user, err := app.models.Users.GetForToken(input.TokenPlaintext, tokens.ScopeActivation)
	if err != nil {
		switch {
		case errors.Is(err, internal.ErrNoRecord):
			v.AddError("token", "invalid or expired activation token")
			app.ValidationErrorResponse(w, r, &v)
		default:
			app.ServerErrorResponse(w, r)
		}
		return
	}

	user.Activated = true
	err = app.models.Users.UpdateUser(user)
	if err != nil {
		app.logger.LogError.Println(err)
		app.ServerErrorResponse(w, r)
		return
	}

	err = app.models.Tokens.DeleteAllForUser(tokens.ScopeActivation, user.Id)
	if err != nil {
		app.logger.LogError.Println(err)
		app.ServerErrorResponse(w, r)
		return
	}

	if err := app.WriteJson(w, r, envelope{"user": user}, http.Header{}, 200); err != nil {
		app.logger.LogError.Println(err)
		app.ServerErrorResponse(w, r)
		return
	}
}

func (app *Application) CreateAuthTokenHanler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := app.ReadJson(r, &input)
	if err != nil {
		app.logger.LogError.Println(err)
		app.InvalidCredentialsResponse(w, r)
		return
	}

	v := validation.New()
	users.ValidateEmail(&v, input.Email)
	users.ValidatePasswordPlainText(&v, input.Password)

	if !v.Valid() {
		app.ValidationErrorResponse(w, r, &v)
		return
	}

	user, err := app.models.Users.GetUserByEmail(input.Email)
	if err != nil {
		switch {
		case errors.Is(err, internal.ErrNoRecord):
			app.InvalidCredentialsResponse(w, r)
		default:
			app.ServerErrorResponse(w, r)
		}
		app.logger.LogError.Println(err)
		return
	}

	match, err := user.Password.Matches(input.Password)

	if err != nil {
		app.logger.LogError.Println(err)
		app.ServerErrorResponse(w, r)
		return
	}

	if !match {
		app.InvalidCredentialsResponse(w, r)
		return
	}

	accessToken, err := app.models.Tokens.New(user.Id, 24*time.Hour, tokens.ScopeAuthentication)
	if err != nil {
		app.logger.LogError.Println(err)
		app.ServerErrorResponse(w, r)
		return
	}

	err = app.WriteJson(w, r, envelope{"authentication_token": accessToken}, http.Header{}, http.StatusCreated)
	if err != nil {
		app.logger.LogError.Println(err)
		app.ServerErrorResponse(w, r)
	}
}
