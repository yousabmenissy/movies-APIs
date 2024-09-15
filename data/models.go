package data

import (
	"database/sql"
	"movies_api/data/movies"
	"movies_api/data/permissions"
	"movies_api/data/tokens"
	"movies_api/data/users"
)

type Models struct {
	Moives      movies.MovieModel
	Users       users.UserModel
	Tokens      tokens.TokenModel
	Permissions permissions.PermissionsModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Moives:      movies.MovieModel{Db: db},
		Users:       users.UserModel{Db: db},
		Tokens:      tokens.TokenModel{DB: db},
		Permissions: permissions.PermissionsModel{Db: db},
	}
}
