package users

import (
	"crypto/sha256"
	"database/sql"
	"errors"
	"fmt"
	"movies_api/internal"
	"time"
)

func (m *UserModel) GetUserByEmail(email string) (*User, error) {
	statement := `SELECT id, name, email, password_hash, activated, created_at, updated_at FROM users WHERE email = $1`

	var user User
	err := m.Db.QueryRow(statement, email).Scan(&user.Id, &user.Name, &user.Email, &user.Password.Hash, &user.Activated, &user.Created_at, &user.Updated_at)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, internal.ErrNoRecord

		default:
			return nil, err
		}
	}

	return &user, nil
}

func (m *UserModel) GetForToken(tokenPlaintext, scope string) (*User, error) {
	tokenHash := sha256.Sum256([]byte(tokenPlaintext))

	query := `SELECT 
		u.id, 
		u.name, 
		u.email, 
		u.password_hash, 
		u.activated, 
		u.created_at, 
		u.updated_at 
	FROM 
		users u 
	INNER JOIN 
		tokens t ON u.id = t.user_id 
	WHERE 
		t.hash = $1 
		AND t.scope = $2
		AND t.expiry > $3
	LIMIT 1`

	var user User
	err := m.Db.QueryRow(query, tokenHash[:], scope, time.Now()).Scan(
		&user.Id,
		&user.Name,
		&user.Email,
		&user.Password.Hash,
		&user.Activated,
		&user.Created_at,
		&user.Updated_at,
	)

	switch {
	case err == sql.ErrNoRows:
		return nil, internal.ErrNoRecord
	case err != nil:
		return nil, fmt.Errorf("failed to retrieve user: %w", err)
	default:
		return &user, nil
	}
}
