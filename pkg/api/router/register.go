package router

import (
	"net/http"

	"github.com/dwethmar/atami/pkg/api/handler"
	"github.com/dwethmar/atami/pkg/auth"

	"github.com/go-chi/chi"
)

// NewRegisterRouter returns the api routes handler
func NewRegisterRouter(service *auth.Service) http.Handler {
	r := chi.NewRouter()

	r.Post("/", handler.Register(service))

	return r
}
