package log

import (
	"StockbitGolang/Microservices/utils/times"
	"database/sql"
)

type Repository struct {
	postgres *sql.DB
}

func NewRepository(postgres *sql.DB) *Repository  {
	return &Repository{postgres: postgres}
}

func (r *Repository) InsertLog(log SearchLog) (err error)  {
	now := times.Now(times.TimeGmt, times.TimeFormat)

	query := `INSERT INTO search_log(id, data, created_at) VALUES($1, $2, $3)`

	_, err = r.postgres.Exec(query, log.ID, log.Data, now)

	return err
}
