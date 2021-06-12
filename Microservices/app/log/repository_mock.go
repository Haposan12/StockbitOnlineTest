package log

import "github.com/stretchr/testify/mock"

type MockRepository struct {
	mock.Mock
}

func (repo *MockRepository) InsertLog(log SearchLog) (err error)  {
	args := repo.Called(log)
	return args.Error(0)
}
