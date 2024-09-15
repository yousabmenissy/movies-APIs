package permissions

import "github.com/lib/pq"

func (m PermissionsModel) AddForUser(user_id int64, codes ...string) error {
	query := `INSERT INTO users_permissions 
	SELECT $1, permissions.id FROM permissions WHERE permissions.code = ANY($2)`

	_, err := m.Db.Exec(query, user_id, pq.Array(codes))

	return err
}
