INSERT INTO movies (
        title,
        director,
        producers,
        prod_companies,
        writers,
        overview,
        status,
        budget,
        age_rating,
        language,
        runtime,
        cast_members,
        genres,
        release_date,
        country
    )
VALUES (
        'Django Unchained',
        'Quentin Tarantino',
        Array ['Pilar Savone', 'Stacey Sher'],
        Array ['The Weinstein Company'],
        Array ['Quentin Tarantino'],
        'With the help of a German bounty-hunter, a freed slave sets out to rescue his wife from a brutal plantation owner in Mississippi.',
        'released',
        100000000,
        'R',
        'english',
        165,
        Array ['Jamie Foxx', 'Christoph Waltz', 'Leonardo DiCaprio', 'Kerry Washington', 'Samuel L. Jackson', 'Walton Goggins'],
        Array ['Drama', 'Western'],
        '2012-12-25',
        'USA'
    )