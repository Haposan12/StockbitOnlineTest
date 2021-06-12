package movie

import (
	"StockbitGolang/Microservices/utils/response"
	"net/http"
)

type HTTPDelivery struct {
	MovieService IService
}

func NewDelivery(MovieSearch IService) *HTTPDelivery  {
	return &HTTPDelivery{MovieService: MovieSearch}
}

func (d *HTTPDelivery) GetMoviesSearch(w http.ResponseWriter, r *http.Request) {
	var err error

	//get param
	searchWord := r.URL.Query().Get(SearchWordParam)
	pagination := r.URL.Query().Get(PaginationParam)

	//check param
	if searchWord == "" || pagination == "" {
		err = CheckParamError
		response.SendErrorResponse(w, err.Error(), http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	//service
	data, err := d.MovieService.GetMoviesService(searchWord, pagination)
	if err != nil {
		response.SendErrorResponse(w, err.Error(), http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	response.SendSuccessResponse(w, data, http.StatusText(http.StatusOK))
}

func (d *HTTPDelivery)GetMovieDetail(w http.ResponseWriter, r *http.Request) {
	var err error

	//get param
	movieTitle := r.URL.Query().Get(MovieTitleParam)

	if movieTitle == "" {
		err = CheckParamError
		response.SendErrorResponse(w, err.Error(), http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	//service
	data, err := d.MovieService.GetMovieDetailService(movieTitle)
	if err != nil {
		response.SendErrorResponse(w, err.Error(), http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	response.SendSuccessResponse(w, data, http.StatusText(http.StatusOK))
}

