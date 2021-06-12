package log

import "testing"

func TestService_InsertSearchLog(t *testing.T) {
	mockRepository := MockRepository{}

	log := SearchLog{}

	mockRepository.On("InsertLog", log).Return(nil)
}
