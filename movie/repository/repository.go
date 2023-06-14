package repository

import (
	"context"

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

func (r *MovieRepository) FindOne(seriesTitle string) (*domain.Movie, error) {
	sql := "select * from movies where series_title=$1"

	var movie domain.Movie

	err := r.Connection.QueryRow(
		context.Background(), sql, seriesTitle).Scan(
			&movie.PosterLink,
			&movie.SeriesTitle, 
			&movie.ReleasedYear, 
			&movie.Certificate, 
			&movie.Runtime, 
			&movie.Genre, 
			&movie.ImdbRating, 
			&movie.Overview, 
			&movie.MetaScore, 
			&movie.Director, 
			&movie.Start1, 
			&movie.Start2, 
			&movie.Start3, 
			&movie.Start4, 
			&movie.NoOfVotes, 
			&movie.Gross)

	if err != nil {
		return nil, err
	}

	return &movie, nil
}

func (r *MovieRepository) FindAll() ([]*domain.Movie, error) {
	// TODO
	return nil, nil
}