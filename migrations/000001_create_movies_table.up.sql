CREATE TABLE IF NOT EXISTS movies (
    id BIGSERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    director TEXT NOT NULL,
    producers TEXT [] NOT NULL,
    prod_companies TEXT [] NOT NULL,
    writers TEXT [] NOT NULL,
    overview TEXT NOT NULL,
    status TEXT NOT NULL,
    budget INT NOT NULL CHECK (budget > 0),
    age_rating TEXT NOT NULL,
    language TEXT NOT NULL,
    runtime INT NOT NULL CHECK (runtime > 0),
    cast_members TEXT [] NOT NULL,
    genres TEXT [] NOT NULL CHECK (
        array_length(genres, 1) BETWEEN 1 AND 5
    ),
    release_date DATE NOT NULL CHECK (
        release_date >= '1888-01-01'::DATE
        AND release_date <= CURRENT_DATE
    ),
    country TEXT NOT NULL,
    created_at TIMESTAMP(0) WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP(0) WITH TIME ZONE DEFAULT NOW()
);