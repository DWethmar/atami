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

func doNothing(w http.ResponseWriter, r *http.Request) {}

// NewAPI create new API.
func NewAPI(handler http.Handler) *API {
	s := &API{
		router: chi.NewRouter(),
	}

	s.router.Get("/favicon.ico", doNothing)
	s.router.Mount("/", handler)

	return s
}
