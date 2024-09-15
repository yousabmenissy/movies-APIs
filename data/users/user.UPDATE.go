package users

import (
	"database/sql"
	"errors"
	"movies_api/internal"
)

var ErrEditConflict = errors.New("conflicting update statement")

func (m *UserModel) UpdateUser(user *User) error {
	statement := `UPDATE users SET ("name", "email", "password_hash", "activated") = ($2, $3, $4, $5) WHERE id = $1`
	_, err := m.Db.Exec(statement, user.Id, user.Name, user.Email, user.Password.Hash, user.Activated)
	if err != nil {
		switch {
		case err.Error() == `pq: duplicate key value violates unique constraint "users_email_key"`:
			return internal.ErrDuplicateEmail

		case errors.Is(err, sql.ErrNoRows):
			return ErrEditConflict

		default:
			return err
		}
	}
	return nil
}
