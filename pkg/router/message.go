package router

import (
	"net/http"

	"github.com/dwethmar/atami/pkg/handler"
	"github.com/go-chi/chi"
)

// ContentTypesRoutes returns the api routes handler
func messageRoutes() http.Handler {
	r := chi.NewRouter()

	r.Get("/", handler.ListMessages())

	return r
}
