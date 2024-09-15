package users

import (
	"database/sql"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

var AnnonymousUser = &User{}

type User struct {
	Id         int64     `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Password   password  `json:"-"`
	Activated  bool      `json:"activated"`
	Created_at time.Time `json:"created-at"`
	Updated_at time.Time `json:"updated-at"`
}

type password struct {
	plaintext *string
	Hash      []byte
}

type UserModel struct {
	Db *sql.DB
}

func (p *password) Set(passwordPlainText string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(passwordPlainText), 12)
	if err != nil {
		return err
	}

	p.plaintext = &passwordPlainText
	p.Hash = hash

	return nil
}

func (p *password) Matches(passwordPlainText string) (bool, error) {
	err := bcrypt.CompareHashAndPassword(p.Hash, []byte(passwordPlainText))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return false, nil

		default:
			return false, err
		}
	}

	return true, nil
}

func (u *User) IsAnonymous() bool {
	return u == AnnonymousUser
}
