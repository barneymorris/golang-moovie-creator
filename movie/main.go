package main

import (
	"log"

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

	movieController := controller.NewMovieController(movieService, movieRepository)

	app := fiber.New()

	api := app.Group("/api")

	api.Get("/movies/:id", movieController.GetMovieById)

	log.Fatal(app.Listen(":9000"))
}