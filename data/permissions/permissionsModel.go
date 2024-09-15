package permissions

import "database/sql"

type Permissions []string

func (p Permissions) Include(code string) bool {
	for i := range p {
		if code == p[i] {
			return true
		}
	}
	return false
}

type PermissionsModel struct {
	Db *sql.DB
}
