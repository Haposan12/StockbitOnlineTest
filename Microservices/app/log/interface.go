package log

type IService interface {
	InsertSearchLog(SearchLog) error
}

type IRepository interface {
	InsertLog(SearchLog) error
}
