package movie

import "StockbitGolang/Microservices/infrastructure/omdb"

type Movies struct {
	Movies      []omdb.OmdbMovieSearch `json:"movies"`
	TotalResult string            `json:"total_result"`
}

