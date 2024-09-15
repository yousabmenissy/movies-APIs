package movies

import (
	"movies_api/internal"
)

func (m MovieModel) DeleteMovie(id int64) error {
	ok := m.exists(id)
	if !ok {
		return internal.ErrNoRecord
	}
	statement := `DELETE FROM movies WHERE id = $1`
	_, err := m.Db.Exec(statement, id)
	if err != nil {
		return err
	}
	return nil
}
