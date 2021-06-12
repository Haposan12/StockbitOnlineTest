package config

import (
	"StockbitGolang/Microservices/app/constant"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func InitDB() (*sql.DB, error) {
	connInfo := fmt.Sprintf(`postgres://%s:%s@%s:%s/%s?sslmode=disable`, ENV[constant.POSTGRE_USERNAME], ENV[constant.POSTGRE_PASSWORD],
			ENV[constant.POSTGRE_ADDRESS], ENV[constant.POSTGRE_PORT], ENV[constant.POSTGRE_DB])

	db, err := sql.Open("postgres", connInfo)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return db, nil
}
