package main

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/betelgeusexru/golang-moovie-creator/movie/controller"
	"github.com/betelgeusexru/golang-moovie-creator/movie/db"
	"github.com/betelgeusexru/golang-moovie-creator/movie/repository"
	"github.com/betelgeusexru/golang-moovie-creator/movie/service"
	"github.com/gofiber/fiber/v2"
)

func main() {
	moviePostgresClient := db.GetNewPostgresClient()
	movieRepository := repository.NewMovieService(moviePostgresClient)
	movieService := service.NewMovieService(movieRepository)

	defer func(){
		moviePostgresClient.Connection.Close(context.Background())
	}()

	movieController := controller.NewMovieController(movieService, movieRepository)

	app := fiber.New(
		fiber.Config{
			ErrorHandler: func(ctx *fiber.Ctx, err error) error {
				code := fiber.StatusInternalServerError
		
				var e *fiber.Error
				if errors.As(err, &e) {
					code = e.Code
				}

				err = ctx.Status(code).JSON(map[string]string{"error": e.Error()})

				if err != nil {
					return ctx.Status(http.StatusInternalServerError).JSON(map[string]string{"error": "something went wrong, plesae try later"})
				}
		
				return nil
			},
		},
	)

	api := app.Group("/api")

	api.Get("/movies/:id", movieController.GetMovieById)
	api.Get("/movies", movieController.GetAllMovies)

	log.Fatal(app.Listen(":9000"))
}