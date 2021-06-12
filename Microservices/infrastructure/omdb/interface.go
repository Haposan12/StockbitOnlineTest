package omdb

type IService interface {
	GetMovies(string, string) (OmdbMovies, error)
	GetMovie(string) (MovieDetail, error)
}
