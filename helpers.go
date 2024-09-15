package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"movies_api/internal"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type envelope map[string]interface{}

func (config *Config) OpenConnection() (*sql.DB, error) {
	connString := fmt.Sprintf(`host=%s port=%s dbname=%s user=%s password=%s sslmode=disable`,
		config.DB.DbHost, config.DB.Dbport, config.DB.Dbname, config.DB.Dbuser, config.DB.Dbpassword,
	)
	dbConn, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}

	dbConn.SetMaxOpenConns(100)
	dbConn.SetMaxIdleConns(100)
	dbConn.SetConnMaxIdleTime(time.Hour)
	return dbConn, nil
}

func (app *Application) WriteJson(w http.ResponseWriter, r *http.Request, data envelope, header http.Header, status int) error {
	qs := r.URL.Query()
	print := app.ReadString(qs, "print", "")

	if print == "pretty" {
		d, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			return err
		}
		for key, val := range header {
			w.Header()[key] = val
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		w.Write(d)

		return nil
	} else {
		d, err := json.Marshal(data)
		if err != nil {
			return err
		}
		for key, val := range header {
			w.Header()[key] = val
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		w.Write(d)
		return nil
	}
}

func (app *Application) ReadJson(r *http.Request, data interface{}) error {
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		return err
	}

	return nil
}

func (app *Application) ReadString(qs url.Values, key string, defaultValue string) string {
	s := qs.Get(key)

	if s == "" {
		return defaultValue
	}

	return s
}

func (app *Application) ReadStrings(qs url.Values, key string, defaultValue []string) []string {
	s := qs.Get(key)

	if s == "" {
		return defaultValue
	}

	return strings.Split(s, ",")
}

func (app *Application) ReadInt(qs url.Values, key string, defaultValue int) int {
	i := qs.Get(key)

	if i == "" {
		return defaultValue
	}

	num, _ := strconv.Atoi(i)
	return num
}

func (app *Application) readURLParams(qs url.Values, params *internal.URLParams) {
	params.Title = app.ReadString(qs, "title", "")
	params.Director = app.ReadString(qs, "director", "")
	params.Producers = app.ReadStrings(qs, "producers", []string{})
	params.Prod_companies = app.ReadStrings(qs, "prod-companies", []string{})
	params.Writers = app.ReadStrings(qs, "writers", []string{})
	params.Cast_members = app.ReadStrings(qs, "cast-members", []string{})
	params.Genres = app.ReadStrings(qs, "genres", []string{})
	params.Age_rating = app.ReadString(qs, "age-rating", "")
	params.Status = app.ReadString(qs, "status", "")
	params.Country = app.ReadString(qs, "country", "")
	params.Sort = app.ReadString(qs, "sort", "id")
	params.Print = app.ReadString(qs, "print", "default")
	params.Page = app.ReadInt(qs, "page", 0)
	params.Page_size = app.ReadInt(qs, "page_size", 0)
}

func (app *Application) background(fn func()) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				app.logger.LogError.Printf("%s", err)
			}
		}()
		fn()
	}()
}
