package users

import "movies_api/internal"

func (m *UserModel) InserUser(user *User) error {
	statement := `INSERT INTO users ("name", "email", "password_hash", "activated") VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at`
	if err := m.Db.QueryRow(statement, user.Name, user.Email, user.Password.Hash, user.Activated).Scan(&user.Id, &user.Created_at, &user.Updated_at); err != nil {
		switch {
		case err.Error() == `pq: duplicate key value violates unique constraint "users_email_key"`:
			return internal.ErrDuplicateEmail
		}
		return err
	}

	return nil
}
