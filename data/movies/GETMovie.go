package movies

import (
	"fmt"
	"movies_api/internal"
	"time"

	"github.com/lib/pq"
)

func (m MovieModel) GetMovies(params internal.URLParams) ([]Movie, error) {
	statement :=
		fmt.Sprintf(`SELECT "id", "title", "director", "producers", "prod_companies", "writers", "overview", "status", 
	"budget", "age_rating", "language", "runtime", "cast_members", "genres", "release_date", "country", "created_at", "updated_at" FROM movies
	WHERE (to_tsvector('english', title) @@ plainto_tsquery('english', $1) OR $1 = '')
	AND (to_tsvector('english', director) @@ plainto_tsquery('english', $2) OR $2 = '')
	AND (producers @> $3 OR $3 = '{}')
	AND (prod_companies @> $4 OR $4 = '{}')
	AND (writers @> $5 OR $5 = '{}')
	AND (cast_members @> $6 OR $6 = '{}')
	AND (to_tsvector('english', age_rating) @@ plainto_tsquery('english', $7) OR $7 = '')
	AND (to_tsvector('english', status) @@ plainto_tsquery('english', $8) OR $8 = '')
	AND (to_tsvector('english', country) @@ plainto_tsquery('english', $9) OR $9 = '')
	AND (genres @> $10 OR $10 = '{}')
	ORDER BY %s %s, id ASC
	LIMIT %s OFFSET $11`, params.SortColumn(), params.SortDirection(), params.Limit())

	rows, err := m.Db.Query(
		statement,
		params.Title,
		params.Director,
		pq.Array(params.Producers),
		pq.Array(params.Prod_companies),
		pq.Array(params.Writers),
		pq.Array(params.Cast_members),
		params.Age_rating,
		params.Status,
		params.Country,
		pq.Array(params.Genres),
		params.Offset(),
	)

	var allBooks []Movie

	if err != nil {
		return allBooks, err
	}

	for rows.Next() {
		var newMovie Movie
		var (
			release_date time.Time
			created_at   time.Time
			updated_at   time.Time
		)

		err := rows.Scan(
			&newMovie.Id,
			&newMovie.Title,
			&newMovie.Director,
			pq.Array(&newMovie.Producers),
			pq.Array(&newMovie.Prod_companies),
			pq.Array(&newMovie.Writers),
			&newMovie.Overview,
			&newMovie.Status,
			&newMovie.Budget,
			&newMovie.Age_rating,
			&newMovie.Language,
			&newMovie.Runtime,
			pq.Array(&newMovie.Cast_members),
			pq.Array(&newMovie.Genres),
			&release_date,
			&newMovie.Country,
			&created_at,
			&updated_at,
		)
		if err != nil {
			return []Movie{}, err
		}
		newMovie.Release_date = fmt.Sprint(release_date.Month(), "-", release_date.Day(), "-", release_date.Year())
		newMovie.Created_at = fmt.Sprint(created_at.Month(), "-", created_at.Day(), "-", created_at.Year())
		newMovie.Updated_at = fmt.Sprint(updated_at.Month(), "-", updated_at.Day(), "-", updated_at.Year(), " ", updated_at.Hour(), ":", updated_at.Minute(), ":", updated_at.Second())
		allBooks = append(allBooks, newMovie)
	}

	defer rows.Close()
	return allBooks, nil
}

func (m MovieModel) GetMovie(id int) (*Movie, error) {
	var (
		newMovie     Movie
		release_date time.Time
		created_at   time.Time
		updated_at   time.Time
	)

	statement := `SELECT 
	"id", "title", "director", "producers", "prod_companies", "writers", "overview", "status", 
	"budget", "age_rating", "language", "runtime", "cast_members", "genres", "release_date", "country", "created_at", "updated_at"
FROM movies WHERE id = $1`

	err := m.Db.QueryRow(statement, id).Scan(
		&newMovie.Id,
		&newMovie.Title,
		&newMovie.Director,
		pq.Array(&newMovie.Producers),
		pq.Array(&newMovie.Prod_companies),
		pq.Array(&newMovie.Writers),
		&newMovie.Overview,
		&newMovie.Status,
		&newMovie.Budget,
		&newMovie.Age_rating,
		&newMovie.Language,
		&newMovie.Runtime,
		pq.Array(&newMovie.Cast_members),
		pq.Array(&newMovie.Genres),
		&release_date,
		&newMovie.Country,
		&created_at,
		&updated_at,
	)
	if err != nil {
		return &Movie{}, err
	}

	newMovie.Release_date = fmt.Sprint(release_date.Month(), "-", release_date.Day(), "-", release_date.Year())
	newMovie.Created_at = fmt.Sprint(created_at.Month(), "-", created_at.Day(), "-", created_at.Year())
	newMovie.Updated_at = fmt.Sprint(updated_at.Month(), "-", updated_at.Day(), "-", updated_at.Year(), " ", updated_at.Hour(), ":", updated_at.Minute(), ":", updated_at.Second())
	return &newMovie, nil
}
