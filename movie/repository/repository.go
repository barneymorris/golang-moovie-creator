package repository

import (
	"context"
	"fmt"

	"github.com/betelgeusexru/golang-moovie-creator/movie/db"
	"github.com/betelgeusexru/golang-moovie-creator/movie/domain"
	"github.com/sirupsen/logrus"
)

type Repository interface {
	FindOne(id string) (*domain.Movie, error)
	FindAll() (*[]domain.Movie, error)
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
			&movie.Gross);
	
	var loggingPrams map[string]any = make(map[string]any)
	loggingPrams["seriesTitle"] = seriesTitle

	logrus.WithFields(logrus.Fields{
		"sql": sql,
		"method": "movie repository :: FindOne",
		"params": loggingPrams,
	}).Info("repository fetch")
		

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"sql": sql,
			"method": "movie repository :: FindOne",
			"params": loggingPrams,
			"err": err.Error(),
		}).Warn("repository fetch error")

		return nil, err
	}

	return &movie, nil
}

func (r *MovieRepository) FindAll(limit int, page int) (*[]domain.Movie, error) {
	sql := "select * from movies"

	logrus.WithFields(logrus.Fields{
		"sql": sql,
		"method": "movie repository :: FindAll",
	}).Info("repository fetch")

	var movies []domain.Movie

	rows, err := r.Connection.Query(context.Background(), sql)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"sql": sql,
			"method": "movie repository :: FindAll",
		}).Error("repository fetch error")
	}

	fmt.Printf("page=%d\n", page)
	fmt.Printf("limit=%d\n", limit)

	counter := 0
	for rows.Next()  {
		fmt.Printf("counter=%d\n", counter)
		if counter == (limit * page) {
			break
		}

		var movie domain.Movie
    	err := rows.Scan(&movie.PosterLink,
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
			logrus.WithFields(logrus.Fields{
				"sql": sql,
				"method": "movie repository :: FindAll",
			}).Error("repository fetch error")
		}
   		
		movies = append(movies, movie)
		counter++
	}

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"sql": sql,
			"method": "movie repository :: FindAll",
			"err": err.Error(),
		}).Warn("repository fetch error")

		return nil, err
	}

	defer func(){
		rows.Close()
	}()

	return &movies, nil
}