package router

import (
	"net/http"

	"github.com/go-chi/chi"
)

// NewRouter creates a new api router.
func NewRouter() http.Handler {
	r := chi.NewRouter()

	r.Mount("/messages", messageRoutes())

	return r
}
