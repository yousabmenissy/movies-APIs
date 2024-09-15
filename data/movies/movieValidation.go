package movies

import "movies_api/internal/validation"

func ValidateMovie(v *validation.Validator, movie *Movie) {
	v.Check(movie.Title != "", "title", "must be provided")
	v.Check(movie.Director != "", "director", "must be provided")
	v.Check(movie.Producers != nil, "producers", "must be provided")
	v.Check(movie.Prod_companies != nil, "production comapnies", "must be provided")
	v.Check(movie.Writers != nil, "writers", "must be provided")
	v.Check(movie.Cast_members != nil, "cast members", "must be provided")
	v.Check(movie.Age_rating != "", "age rating", "must be provided")
	v.Check(movie.Status != "", "release status", "must be provided")
	v.Check(movie.Country != "", "country", "must be provided")
	v.Check(movie.Budget > 0, "budget", "must be greater than 0")
	v.Check(v.Unique(movie.Genres), "genres", "duplicates not allowed")
	v.Check(v.Unique(movie.Prod_companies), "production-comapnies", "duplicates not allowed")
	v.Check(v.Unique(movie.Producers), "producers", "duplicates not allowed")
	v.Check(v.Unique(movie.Writers), "writer", "duplicates not allowed")
	v.Check(v.Unique(movie.Cast_members), "cast-members", "duplicates not allowed")
}
