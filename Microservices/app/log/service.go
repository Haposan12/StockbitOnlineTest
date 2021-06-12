package log

type Service struct {
	LogRepository IRepository
}

func NewService(LogRepository IRepository) *Service  {
	return &Service{LogRepository: LogRepository}
}

func (s *Service) InsertSearchLog(log SearchLog) error  {
	return s.LogRepository.InsertLog(log)
}
