package movie

import (
	"StockbitGolang/Microservices/app/log"
	"StockbitGolang/Microservices/infrastructure/omdb"
	"StockbitGolang/Microservices/utils/uuid"
	"encoding/json"
)

type Service struct {
	LogService log.IService
	OmdbService omdb.IService
}

func NewService(LogService log.IService, OmdbService omdb.IService) *Service  {
	return &Service{LogService: LogService, OmdbService: OmdbService}
}

func (s *Service) GetMoviesService(searchWord, pagination string) (movies Movies, err error) {
	omdbMovies, err := s.OmdbService.GetMovies(searchWord, pagination)

	if err != nil {
		return
	}

	//insert log with go routine
	go func() {
		//set log
		logStruct := log.Log{
			SearchType: PaginationSearch,
			SearchWord: searchWord,
		}

		logData, err := json.Marshal(logStruct)
		if err != nil {
			return
		}

		s.LogService.InsertSearchLog(log.SearchLog{
			ID:        uuid.GenerateId(),
			Data:      logData,
		})
	}()

	movies.Movies = omdbMovies.Search
	movies.TotalResult = omdbMovies.TotalResults

	return
}

func (s *Service) GetMovieDetailService(movieTitle string) (omdb.MovieDetail, error) {
	//get movie detail
	movieDetail, err := s.OmdbService.GetMovie(movieTitle)
	if err != nil {
		return movieDetail, err
	}

	//insert log with go routine
	go func() {
		//set log
		logStruct := log.Log{
			SearchType: DetailSearch,
			SearchWord: movieTitle,
		}

		logData, err := json.Marshal(logStruct)
		if err != nil {
			return
		}

		s.LogService.InsertSearchLog(log.SearchLog{
			ID:        uuid.GenerateId(),
			Data:      logData,
		})
	}()

	return movieDetail, nil
}

