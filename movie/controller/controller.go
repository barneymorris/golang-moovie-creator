package controller

import (
	"github.com/betelgeusexru/golang-moovie-creator/movie/domain"
	"github.com/betelgeusexru/golang-moovie-creator/movie/repository"
	"github.com/betelgeusexru/golang-moovie-creator/movie/service"
	"github.com/gofiber/fiber/v2"
)

type Controller interface {
	GetMovieById(id string) (*domain.Movie, error)
}

type MovieController struct {
	*service.MovieService
	*repository.MovieRepository
}

func NewMovieController(svc *service.MovieService, repo *repository.MovieRepository) *MovieController {
	return &MovieController{
		MovieService: svc,
		MovieRepository: repo,
	}
}

func (h *MovieController) GetMovieById(c *fiber.Ctx) error {
	return c.Status(200).JSON("working fine")
}