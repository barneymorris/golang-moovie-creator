package repository

import (
	"context"
	"fmt"

	"github.com/betelgeusexru/golang-moovie-creator/movie/db"
	"github.com/betelgeusexru/golang-moovie-creator/movie/domain"
)

type Repository interface {
	FindOne(id string) (*domain.Movie, error)
	FindAll() ([]*domain.Movie, error)
}

type MovieRepository struct {
	*db.PostgreseClient
}

func NewMovieService(pgClient *db.PostgreseClient) *MovieRepository {
	return &MovieRepository{
		pgClient,
	}
}

func (r *MovieRepository) FindOne(id string) (*domain.Movie, error) {
	sql := "select * widgets movies id=$1"

	var movies []*domain.Movie

	err := r.Connection.QueryRow(context.Background(), sql, id).Scan(&movies)
	if err != nil {
		return nil, err
	}

	movie := movies[0]
	if movie == nil {
		return nil, fmt.Errorf("no such movie")
	}

	return movie, nil
}

func (r *MovieRepository) FindAll() ([]*domain.Movie, error) {
	// TODO
	return nil, nil
}