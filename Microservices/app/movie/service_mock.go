package movie

import (
	"StockbitGolang/Microservices/infrastructure/omdb"
	"github.com/stretchr/testify/mock"
)

type ServiceMock struct {
	mock.Mock
}

func (service *ServiceMock) GetMoviesService(searchWord, pagination string) (movies Movies, err error)  {
	args := service.Called(searchWord, pagination)
	return args.Get(0).(Movies), args.Error(1)
}

func (service *ServiceMock) GetMovieDetailService(movieTitle string) (omdb.MovieDetail, error) {
	args := service.Called(movieTitle)
	return args.Get(0).(omdb.MovieDetail), args.Error(1)
}
