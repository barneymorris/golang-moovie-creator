package controller

import (
	"net/http"

	"github.com/betelgeusexru/golang-moovie-creator/movie/domain"
	"github.com/betelgeusexru/golang-moovie-creator/movie/repository"
	"github.com/betelgeusexru/golang-moovie-creator/movie/service"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
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
	logrus.Info("hit movie controller :: GetMovieByiD")

	id := c.Params("id")

	if id == "" {
		return fiber.NewError(http.StatusBadRequest, "you didnt provide any id")
	}

	movie, err := h.MovieService.FindOne(id)

	if err != nil {
		params := make(map[string]string)
		params["id"] = id

		logrus.WithFields(logrus.Fields{
			"err": err.Error(),
			"params": params,
		}).Warn("movie controller error :: GetMovieByiD")

		return fiber.NewError(http.StatusBadRequest, err.Error())
	}

	return c.Status(200).JSON(movie)
}