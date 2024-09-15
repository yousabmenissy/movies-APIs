package movies

import "database/sql"

type Movie struct {
	Id             int64    `json:"id"`
	Title          string   `json:"title"`
	Director       string   `json:"director"`
	Producers      []string `json:"producers"`
	Prod_companies []string `json:"prod-companies"`
	Writers        []string `json:"writers"`
	Overview       string   `json:"overview"`
	Status         string   `json:"status"`
	Budget         int      `json:"budget"`
	Age_rating     string   `json:"age-rating"`
	Language       string   `json:"language"`
	Runtime        int      `json:"runtime"`
	Cast_members   []string `json:"cast-members"`
	Genres         []string `json:"genres"`
	Release_date   string   `json:"release-date"`
	Country        string   `json:"country"`
	Created_at     string   `json:"created-at"`
	Updated_at     string   `json:"updated-at"`
}

type MovieModel struct {
	Db *sql.DB
}

func (m MovieModel) exists(id int64) bool {
	var movieId int64
	statement := `SELECT "id" FROM movies WHERE id = $1`
	if err := m.Db.QueryRow(statement, id).Scan(&movieId); err != nil {
		return false
	}
	if movieId == id {
		return true
	} else {
		return false
	}
}
