package movie

import "StockbitGolang/Microservices/infrastructure/omdb"

type IService interface {
	GetMoviesService(string, string) (Movies, error)
	GetMovieDetailService(string) (omdb.MovieDetail, error)
}
