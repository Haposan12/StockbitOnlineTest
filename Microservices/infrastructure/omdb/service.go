package omdb

import (
	"StockbitGolang/Microservices/app/constant"
	"StockbitGolang/Microservices/config"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Service struct {

}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GetMovies(searchWord, pagination string) (response OmdbMovies, err error) {
	//set url
	url := fmt.Sprintf("%s?apikey=%s&s=%s&page=%s", config.ENV[constant.URL_OMDB], config.ENV[constant.API_KEY], searchWord, pagination)

	//http request to omdb
	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return
	}

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return
	}

	if err := json.Unmarshal([]byte(string(body)), &response); err != nil {
		return response, err
	}

	if response.Response == "False" {
		err = errors.New(response.Error)
		return
	}

	return
}

//GetMovie is a function to get single detail movie
func (s *Service) GetMovie(movieTitle string) (response MovieDetail, err error) {
	//set url
	url := fmt.Sprintf("%s?apikey=%s&t=%s", config.ENV[constant.URL_OMDB], config.ENV[constant.API_KEY], movieTitle)

	//http request to omdb
	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return
	}

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return
	}

	if err := json.Unmarshal([]byte(string(body)), &response); err != nil {
		return response, err
	}

	if response.Response == "False" {
		err = errors.New(response.Error)
		return
	}

	return
}

