package movies

import (
	"github.com/lib/pq"
)

func (m MovieModel) InsertMovie(newMovie *Movie) error {
	statement := `INSERT INTO movies (
        "title",
        "director",
        "producers",
        "prod_companies",
        "writers",
        "overview",
        "status",
        "budget",
        "age_rating",
        "language",
        "runtime",
        "cast_members",
        "genres",
        "release_date",
        "country"
    ) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15)
	 RETURNING id`

	err := m.Db.QueryRow(statement,
		newMovie.Title,
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
	).Scan(&newMovie.Id)
	if err != nil {
		return err
	}

	return nil
}
