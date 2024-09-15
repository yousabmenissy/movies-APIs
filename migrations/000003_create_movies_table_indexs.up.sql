CREATE INDEX IF NOT EXISTS movie_title_index ON movies USING GIN (to_tsvector('english', title));
CREATE INDEX IF NOT EXISTS movie_director_index ON movies USING GIN (to_tsvector('english', director));
CREATE INDEX IF NOT EXISTS movie_status_index ON movies USING GIN (to_tsvector('english', status));
CREATE INDEX IF NOT EXISTS movie_country_index ON movies USING GIN (to_tsvector('english', country));
CREATE INDEX IF NOT EXISTS movie_producers_index ON movies USING GIN (producers);
CREATE INDEX IF NOT EXISTS movie_prod_companies_index ON movies USING GIN (prod_companies);
CREATE INDEX IF NOT EXISTS movie_writers_index ON movies USING GIN (writers);
CREATE INDEX IF NOT EXISTS movie_cast_members_index ON movies USING GIN (cast_members);