struct TMDBMovie {
    1: bool adult,
    2: string backdropPath,
    // let belongsToCollection: NSNull
    3: i32 budget,
    4: list<TMDBMovieGenre> genres,
    5: string homepage,
    6: i32 id,
    7: string imdbID,
    8: string originalLanguage,
    9: string originalTitle,
    10: string overview,
    11: double popularity,
    12: optional string posterPath,
    13: list<TMDBMovieProductionCompany> productionCompanies,
    14: list<TMDBMovieProductionCountry> productionCountries,
    15: string releaseDate,
    16: i32 revenue,
    17: i32 runtime,
    18: list<TMDBMovieSpokenLanguage> spokenLanguages,
    19: string status,
    20: string tagline,
    21: string title,
    22: bool video,
    23: double voteAverage,
    24: i32 voteCount
}

struct TMDBMovieGenre {
    1: i32 id,
    2: string name
}

struct TMDBMovieProductionCompany {
    1: i32 id,
    2: optional string logoPath,
    3: string name,
    4: string originCountry
}

struct TMDBMovieProductionCountry {
    1: string iso3166_1,
    2: string name
}

struct TMDBMovieSpokenLanguage {
    1: string iso639_1,
    2: string name
}

exception TMDBMovieError {
    1: i32 code,
    2: string message
}

struct TMDBGetDetailsRequest {
    1: i32 id,
    2: string apiKey,
    3: optional string language
}

service MovieService {
    TMDBMovie getDetails(1: TMDBGetDetailsRequest req) throws (1:TMDBMovieError error)
}