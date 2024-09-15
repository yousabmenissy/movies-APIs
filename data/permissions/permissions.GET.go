package permissions

func (m PermissionsModel) GetAllForUser(userID int64) (Permissions, error) {
	query := `SELECT permissions.code FROM permissions 
	INNER JOIN users_permissions ON permissions.id = users_permissions.permission_id
	INNER JOIN users ON users_permissions.user_id = users.id
	WHERE users.id = $1`

	rows, err := m.Db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var permissions Permissions
	for rows.Next() {
		var permission string

		err := rows.Scan(&permission)
		if err != nil {
			return nil, err
		}

		permissions = append(permissions, permission)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return permissions, nil

}
