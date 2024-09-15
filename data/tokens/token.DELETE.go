package tokens

func (m TokenModel) DeleteAllForUser(scope string, userID int64) error {
	query := `
		DELETE FROM tokens
		WHERE scope = $1 AND user_id = $2
		`
	_, err := m.DB.Exec(query, scope, userID)
	return err
}
