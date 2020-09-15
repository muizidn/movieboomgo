package main

import (
	"context"
	"fmt"

	"github.com/muizidn/movieboomgo/thrift/gen-go/movie"
)

type MovieHandler struct {
}

func NewMovieHandler() *MovieHandler {
	return &MovieHandler{}
}

func (p *MovieHandler) Get(ctx context.Context, apiKey string, language string) (r *movie.TMDBMovie, err error) {
	return nil, fmt.Errorf("You are sick!")
}
