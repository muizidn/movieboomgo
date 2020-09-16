package tmdb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"go.uber.org/zap"
)

type RespMovieGetDetails struct {
	Adult               bool                    `json:"adult"`
	BackdropPath        string                  `json:"backdrop_path"`
	BelongsToCollection interface{}             `json:"belongs_to_collection"`
	Budget              int64                   `json:"budget"`
	Genres              []RespGenre             `json:"genres"`
	Homepage            string                  `json:"homepage"`
	ID                  int64                   `json:"id"`
	ImdbID              string                  `json:"imdb_id"`
	OriginalLanguage    string                  `json:"original_language"`
	OriginalTitle       string                  `json:"original_title"`
	Overview            string                  `json:"overview"`
	Popularity          float64                 `json:"popularity"`
	PosterPath          interface{}             `json:"poster_path"`
	ProductionCompanies []RespProductionCompany `json:"production_companies"`
	ProductionCountries []RespProductionCountry `json:"production_countries"`
	ReleaseDate         string                  `json:"release_date"`
	Revenue             int64                   `json:"revenue"`
	Runtime             int64                   `json:"runtime"`
	SpokenLanguages     []RespSpokenLanguage    `json:"spoken_languages"`
	Status              string                  `json:"status"`
	Tagline             string                  `json:"tagline"`
	Title               string                  `json:"title"`
	Video               bool                    `json:"video"`
	VoteAverage         float64                 `json:"vote_average"`
	VoteCount           int64                   `json:"vote_count"`
}

type RespGenre struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type RespProductionCompany struct {
	ID            int64   `json:"id"`
	LogoPath      *string `json:"logo_path"`
	Name          string  `json:"name"`
	OriginCountry string  `json:"origin_country"`
}

type RespProductionCountry struct {
	ISO3166_1 string `json:"iso_3166_1"`
	Name      string `json:"name"`
}

type RespSpokenLanguage struct {
	ISO639_1 string `json:"iso_639_1"`
	Name     string `json:"name"`
}

type RespMovieGetDetailsErr struct {
	StatusCode    int    `json:"status_code"`
	StatusMessage string `json:"status_message"`
	Success       bool   `json:"success"`
}

func (r *RespMovieGetDetailsErr) Error() string {
	return r.StatusMessage
}

func MovieGetDetails(apiKey string, lang *string, id int) (*RespMovieGetDetails, *RespMovieGetDetailsErr, error) {

	urlString := fmt.Sprintf("https://api.themoviedb.org/3/movie/%d", id)

	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()
	sugar.Infow("tmdb_MovieGetDetails",
		"urlString", urlString,
		"attempt", 3,
		"backoff", time.Second,
	)

	client := &http.Client{}

	finalURL, err := url.Parse(urlString)
	if err != nil {
		sugar.Errorf("Failure", err)
	}
	q := finalURL.Query()
	q.Set("api_key", apiKey)
	if lang != nil {
		q.Set("language", *lang)
	}
	finalURL.RawQuery = q.Encode()
	resp, err := client.Get(finalURL.String())
	if err != nil {
		sugar.Errorf("Failure", err)
	}

	sugar.Infof("Response statusCode %d", resp.StatusCode)

	respBody, _ := ioutil.ReadAll(resp.Body)

	switch resp.StatusCode {
	case 200:
		var result RespMovieGetDetails
		json.Unmarshal(respBody, &result)
		return &result, nil, nil
	case 401:
		var result RespMovieGetDetailsErr
		json.Unmarshal(respBody, &result)
		return nil, &result, nil
	default:
		return nil, nil, fmt.Errorf("TMDB StatusCode: %d", resp.StatusCode)
	}
}
