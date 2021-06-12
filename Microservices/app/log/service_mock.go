package log

import "github.com/stretchr/testify/mock"

type ServiceMock struct {
	mock.Mock
}

func (service *ServiceMock) InsertSearchLog(log SearchLog) error  {
	args := service.Called(log)
	return args.Error(0)
}
