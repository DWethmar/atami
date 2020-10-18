package api

import (
	"net/http"

	"github.com/go-chi/chi"
)

// API defines the api type.
type API struct {
	router *chi.Mux
}

func (s *API) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	s.router.ServeHTTP(res, req)
}

// NewAPI create new API.
func NewAPI(router http.Handler) *API {
	s := &API{
		router: chi.NewRouter(),
	}

	s.router.Mount("/", router)

	return s
}
