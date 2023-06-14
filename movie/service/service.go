package service

import (
	"github.com/betelgeusexru/golang-moovie-creator/movie/domain"
	"github.com/betelgeusexru/golang-moovie-creator/movie/repository"
)

type Filter struct {
	Year *int
	SeriesTitle *string
	ReleasedYear *string
	Genre *string
	ImdbRating *float64
	Overview *string
}

type Service interface {
	FindOne(id string) (*domain.Movie, error)
	FindAll() ([]*domain.Movie, error)
}

type MovieService struct {
	*repository.MovieRepository
}

func NewMovieService(repository *repository.MovieRepository) *MovieService {
	return &MovieService{
		repository,
	}
}

func (s *MovieService) FindOne(id string) (*domain.Movie, error) {
	// TODO
	return nil, nil
}

func (s *MovieService) FindAll() ([]*domain.Movie, error) {
	// TODO
	return nil, nil
}