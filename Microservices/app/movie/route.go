package movie

import "github.com/go-chi/chi"

type Route struct {
	Delivery *HTTPDelivery
}

func NewRoute(delivery *HTTPDelivery) *Route  {
	return &Route{Delivery: delivery}
}

func (routes Route) RegisterRoutes(r chi.Router) {
	r.Group(func(r chi.Router) {
		r.Route("/movie", func(r chi.Router) {
			r.Get("/", routes.Delivery.GetMoviesSearch)
			r.Get("/detail", routes.Delivery.GetMovieDetail)
		})
	})
}
