package domain

type Movie struct {
	PosterLink string `json:"poster_link"`
	SeriesTitle string `json:"series_title"`
	ReleasedYear string `json:"released_year"`
	Certificate string `json:"certificate"`
	Runtime string `json:"runtime"`
	Genre string `json:"genre"`
	ImdbRating float64 `json:"imdb_rating"`
	Overview string `json:"overview"`
	MetaScore int64 `json:"meta_score"`
	Director string `json:"director"`
	Start1 string `json:"start1"`
	Start2 string `json:"start2"`
	Start3 string `json:"start3"`
	Start4 string `json:"start4"`
	NoOfVotes string `json:"no_of_votes"`
	Gross string `json:"gross"`
}