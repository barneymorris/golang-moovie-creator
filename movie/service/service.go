package service

import (
	"github.com/betelgeusexru/golang-moovie-creator/movie/domain"
	"github.com/betelgeusexru/golang-moovie-creator/movie/repository"
	"github.com/sirupsen/logrus"
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
	FindOne(seriesTitle string) (*domain.Movie, error)
	FindAll() (*[]domain.Movie, error)
}

type MovieService struct {
	*repository.MovieRepository
}

func NewMovieService(repository *repository.MovieRepository) *MovieService {
	return &MovieService{
		repository,
	}
}

func (s *MovieService) FindOne(seriesTitle string) (*domain.Movie, error) {
	var loggingPrams map[string]any = make(map[string]any)
	loggingPrams["seriesTitle"] = seriesTitle

	logrus.WithFields(logrus.Fields{
		"method": "movie service :: FindOne",
		"params": loggingPrams,
	}).Info("service hit")

	return s.MovieRepository.FindOne(seriesTitle)
}

func (s *MovieService) FindAll(limit int, page int) (*[]domain.Movie, error) {
	logrus.WithFields(logrus.Fields{
		"method": "movie service :: FindAll",
	}).Info("service hit")

	return s.MovieRepository.FindAll(limit, page)
}