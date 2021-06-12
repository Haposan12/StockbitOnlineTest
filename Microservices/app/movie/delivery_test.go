package movie

import (
	"StockbitGolang/Microservices/infrastructure/omdb"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewDelivery(t *testing.T) {
	movieServiceMock := new(Service)

	expected := &HTTPDelivery{movieServiceMock}

	actual := NewDelivery(movieServiceMock)

	assert.Equal(t, expected, actual)
}

func TestHTTPDelivery_GetMoviesSearch(t *testing.T) {
	movieServiceMock := new(ServiceMock)

	searchWord := "Batman"
	pagination := "1"

	url := fmt.Sprintf("/movie?search=%s&pagination=%s", searchWord, pagination)

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		t.Fatal(err)
	}

	//movie data
	movieData := Movies{
		Movies:      nil,
		TotalResult: "",
	}

	resp := httptest.NewRecorder()

	movieServiceMock.On("GetMoviesService", searchWord, pagination).Return(movieData, nil)

	delivery := NewDelivery(movieServiceMock)

	handler := http.HandlerFunc(delivery.GetMoviesSearch)
	handler.ServeHTTP(resp, req)

	// Assert
	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestHTTPDelivery_GetMoviesSearchBadRequest(t *testing.T) {
	movieServiceMock := new(ServiceMock)

	url := fmt.Sprintf("/movie")

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		t.Fatal(err)
	}

	resp := httptest.NewRecorder()

	delivery := NewDelivery(movieServiceMock)

	handler := http.HandlerFunc(delivery.GetMoviesSearch)
	handler.ServeHTTP(resp, req)

	// Assert
	assert.Equal(t, http.StatusBadRequest, resp.Code)
}

func TestHTTPDelivery_GetMoviesSearchErrorGetMoviesService(t *testing.T) {
	movieServiceMock := new(ServiceMock)

	searchWord := "Batman"
	pagination := "1"

	url := fmt.Sprintf("/movie?search=%s&pagination=%s", searchWord, pagination)

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		t.Fatal(err)
	}

	//movie data
	movieData := Movies{
		Movies:      nil,
		TotalResult: "",
	}

	resp := httptest.NewRecorder()

	movieServiceMock.On("GetMoviesService", searchWord, pagination).Return(movieData, errors.New("test error"))

	delivery := NewDelivery(movieServiceMock)

	handler := http.HandlerFunc(delivery.GetMoviesSearch)
	handler.ServeHTTP(resp, req)

	// Assert
	assert.Equal(t, http.StatusInternalServerError, resp.Code)
}

func TestHTTPDelivery_GetMovieDetail(t *testing.T) {
	movieServiceMock := new(ServiceMock)

	movieTitle := "Batman"

	url := fmt.Sprintf("/movie/detail?title=%s", movieTitle)

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		t.Fatal(err)
	}

	//movie data
	movieData := omdb.MovieDetail{
		Title:      "",
		Year:       "",
		Rated:      "",
		Released:   "",
		Runtime:    "",
		Genre:      "",
		Director:   "",
		Writer:     "",
		Actors:     "",
		Plot:       "",
		Language:   "",
		Country:    "",
		Awards:     "",
		Poster:     "",
		Ratings:    nil,
		Metascore:  "",
		ImdbRating: "",
		ImdbVotes:  "",
		ImdbId:     "",
		Type:       "",
		DVD:        "",
		BoxOffice:  "",
		Production: "",
		Website:    "",
		Response:   "",
		Error:      "",
	}

	resp := httptest.NewRecorder()

	movieServiceMock.On("GetMovieDetailService", movieTitle).Return(movieData, nil)

	delivery := NewDelivery(movieServiceMock)

	handler := http.HandlerFunc(delivery.GetMovieDetail)
	handler.ServeHTTP(resp, req)

	// Assert
	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestHTTPDelivery_GetMovieDetailErrorBadRequest(t *testing.T) {
	movieServiceMock := new(ServiceMock)

	movieTitle := ""

	url := fmt.Sprintf("/movie/detail?title=%s", movieTitle)

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		t.Fatal(err)
	}


	resp := httptest.NewRecorder()

	delivery := NewDelivery(movieServiceMock)

	handler := http.HandlerFunc(delivery.GetMovieDetail)
	handler.ServeHTTP(resp, req)

	// Assert
	assert.Equal(t, http.StatusBadRequest, resp.Code)
}

func TestHTTPDelivery_GetMovieDetailErrorGetMovieDetail(t *testing.T) {
	movieServiceMock := new(ServiceMock)

	movieTitle := "Batman"

	url := fmt.Sprintf("/movie/detail?title=%s", movieTitle)

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		t.Fatal(err)
	}

	//movie data
	movieData := omdb.MovieDetail{
		Title:      "",
		Year:       "",
		Rated:      "",
		Released:   "",
		Runtime:    "",
		Genre:      "",
		Director:   "",
		Writer:     "",
		Actors:     "",
		Plot:       "",
		Language:   "",
		Country:    "",
		Awards:     "",
		Poster:     "",
		Ratings:    nil,
		Metascore:  "",
		ImdbRating: "",
		ImdbVotes:  "",
		ImdbId:     "",
		Type:       "",
		DVD:        "",
		BoxOffice:  "",
		Production: "",
		Website:    "",
		Response:   "",
		Error:      "",
	}

	resp := httptest.NewRecorder()

	movieServiceMock.On("GetMovieDetailService", movieTitle).Return(movieData, errors.New("test error"))

	delivery := NewDelivery(movieServiceMock)

	handler := http.HandlerFunc(delivery.GetMovieDetail)
	handler.ServeHTTP(resp, req)

	// Assert
	assert.Equal(t, http.StatusInternalServerError, resp.Code)
}