package main

import (
	"encoding/json"
	"context"

	"github.com/muizidn/movieboomgo/thrift/gen-go/movie"
	"github.com/muizidn/movieboomgo/tmdb"
)

type MovieHandler struct {
}

func NewMovieHandler() *MovieHandler {
	return &MovieHandler{}
}

func (p *MovieHandler) GetDetails(ctx context.Context, req *movie.TMDBGetDetailsRequest) (r *movie.TMDBMovie, err error) {
	var lang string
	if req.IsSetLanguage() {
		lang = req.GetLanguage()
	}
	resp, resperr, syserr := tmdb.MovieGetDetails(req.GetApiKey(), &lang, int(req.GetID()))
	if syserr != nil {
		return nil, syserr
	}
	if resperr != nil {
		return nil, resperr
	}
	jsondata, err := json.Marshal(resp)
	if err != nil {
		return nil, err
	}
	var thriftresp movie.TMDBMovie
	err = json.Unmarshal(jsondata, &thriftresp)
	if err != nil {
		return nil, err
	}
	return &thriftresp, nil
}
