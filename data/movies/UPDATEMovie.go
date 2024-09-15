package movies

import (
	"errors"
	"strconv"

	"github.com/lib/pq"
)

func (m MovieModel) UpdateMovie(newMovie Movie) error {
	if ok := m.exists(newMovie.Id); !ok {
		return errors.New("the movie with the id " + strconv.Itoa(int(newMovie.Id)) + " doesn't exist. and therefore can not be updated")
	}

	statement := `UPDATE movies SET ("title", "director", "producers", "prod_companies", "writers", "overview", "status", 
	"budget", "age_rating", "language", "runtime", "cast_members", "genres", "release_date", "country") = ($2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16)
	WHERE id = $1`

	_, err := m.Db.Exec(statement,
		newMovie.Id, newMovie.Title,
		newMovie.Director,
		pq.Array(newMovie.Producers),
		pq.Array(newMovie.Prod_companies),
		pq.Array(newMovie.Writers),
		newMovie.Overview,
		newMovie.Status,
		newMovie.Budget,
		newMovie.Age_rating,
		newMovie.Language,
		newMovie.Runtime,
		pq.Array(newMovie.Cast_members),
		pq.Array(newMovie.Genres),
		newMovie.Release_date,
		newMovie.Country,
	)
	if err != nil {
		return err
	}
	return nil
}
