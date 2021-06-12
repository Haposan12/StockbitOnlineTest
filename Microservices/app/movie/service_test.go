package movie

import (
	"StockbitGolang/Microservices/app/log"
	"StockbitGolang/Microservices/infrastructure/omdb"
	"StockbitGolang/Microservices/utils/uuid"
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewService(t *testing.T) {
	logServiceMock := new(log.ServiceMock)
	omdbServiceMock := new(omdb.ServiceMock)

	expected := &Service{logServiceMock, omdbServiceMock}

	actual := NewService(logServiceMock, omdbServiceMock)

	assert.Equal(t, expected, actual)
}

func TestService_GetMoviesService(t *testing.T) {
	logServiceMock := new(log.ServiceMock)
	omdbServiceMock := new(omdb.ServiceMock)

	omdbMovies := omdb.OmdbMovies{}

	//param mock
	searchWord := "Batman"
	pagination := "1"

	//log data
	logStruct := log.Log{
		SearchType: PaginationSearch,
		SearchWord: searchWord,
	}

	logData, _ := json.Marshal(logStruct)

	logSearch := log.SearchLog{
		ID:        uuid.GenerateId(),
		Data:      logData,
	}

	omdbServiceMock.On("GetMovies", searchWord, pagination).Return(omdbMovies, nil)

	logServiceMock.On("InsertSearchLog", logSearch).Return(nil)

	//response
	movies := Movies{}

	movies.Movies = omdbMovies.Search
	movies.TotalResult = omdbMovies.TotalResults

	service := NewService(logServiceMock, omdbServiceMock)

	_, err := service.GetMoviesService(searchWord, pagination)

	assert.NoError(t, err)
}

func TestService_GetMoviesServiceErrorWhenGetMovies(t *testing.T) {
	logServiceMock := new(log.ServiceMock)
	omdbServiceMock := new(omdb.ServiceMock)

	omdbMovies := omdb.OmdbMovies{}

	//param mock
	searchWord := "Batman"
	pagination := "1"

	omdbServiceMock.On("GetMovies", searchWord, pagination).Return(omdbMovies, errors.New("test error"))

	service := NewService(logServiceMock, omdbServiceMock)

	_, err := service.GetMoviesService(searchWord, pagination)
	assert.Equal(t, errors.New("test error"), err)
}

func TestService_GetMovieDetailService(t *testing.T) {
	logServiceMock := new(log.ServiceMock)
	omdbServiceMock := new(omdb.ServiceMock)

	movieDetail := omdb.MovieDetail{}

	//log data
	logStruct := log.Log{
		SearchType: PaginationSearch,
		SearchWord: "Batman",
	}

	logData, _ := json.Marshal(logStruct)

	logSearch := log.SearchLog{
		ID:        uuid.GenerateId(),
		Data:      logData,
	}

	omdbServiceMock.On("GetMovie", "Batman").Return(movieDetail, nil)

	logServiceMock.On("InsertSearchLog", logSearch).Return(nil)

	service := NewService(logServiceMock, omdbServiceMock)

	_, err := service.GetMovieDetailService("Batman")

	assert.NoError(t, err)
}

func TestService_GetMovieDetailServiceErrorWhenGetMovies(t *testing.T) {
	logServiceMock := new(log.ServiceMock)
	omdbServiceMock := new(omdb.ServiceMock)

	movieDetail := omdb.MovieDetail{}

	//param mock
	searchWord := "Batman"

	omdbServiceMock.On("GetMovie", searchWord).Return(movieDetail, errors.New("test error"))

	service := NewService(logServiceMock, omdbServiceMock)

	_, err := service.GetMovieDetailService(searchWord)
	assert.Equal(t, errors.New("test error"), err)
}
