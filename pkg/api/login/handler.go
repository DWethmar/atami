package login

import (
	"net/http"

	"github.com/dwethmar/atami/pkg/auth"
	"github.com/go-chi/chi"
	"github.com/go-chi/httplog"
)

// NewHandler returns the api routes handler
func NewHandler(service auth.Service) http.Handler {
	r := chi.NewRouter()

	logger := httplog.NewLogger("login", httplog.Options{})
	r.Use(httplog.RequestLogger(logger))

	r.Post("/", Login(service))

	return r
}
