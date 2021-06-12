package main

import (
	"StockbitGolang/Microservices/app/constant"
	"StockbitGolang/Microservices/app/log"
	"StockbitGolang/Microservices/app/movie"
	"StockbitGolang/Microservices/config"
	"StockbitGolang/Microservices/infrastructure/omdb"
	"StockbitGolang/Microservices/utils/server"
	"database/sql"
	"fmt"
	"github.com/go-chi/chi"
)


var GlobalDbSQL *sql.DB

func main()  {
	var err error

	//init env
	if err = config.LoadEnv(); err != nil {
		panic(err)
	}

	//init db
	GlobalDbSQL, err = config.InitDB()

	/*module*/

	//omdb
	omdbService := omdb.NewService()

	//log
	logRepository := log.NewRepository(GlobalDbSQL)
	logService := log.NewService(logRepository)

	//movie
	movieService := movie.NewService(logService, omdbService)
	movieController := movie.NewDelivery(movieService)
	movieRoutes := movie.NewRoute(movieController)

	//router
	r := chi.NewRouter()
	movieRoutes.RegisterRoutes(r)

	//set port
	port := fmt.Sprintf("%s", config.ENV[constant.PORT])

	fmt.Println("Server is listening to port", port)

	server.Listen(":"+port, r)
}