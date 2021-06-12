package omdb

import "github.com/stretchr/testify/mock"

type ServiceMock struct {
	mock.Mock
}

func (service *ServiceMock) GetMovies(searchWord, pagination string) (response OmdbMovies, err error)  {
	args := service.Called(searchWord, pagination)
	return args.Get(0).(OmdbMovies), args.Error(1)
}

func (service *ServiceMock) GetMovie(movieTitle string) (response MovieDetail, err error)  {
	args := service.Called(movieTitle)
	return args.Get(0).(MovieDetail), args.Error(1)
}
