package login

import (
	"net/http"

	"github.com/dwethmar/atami/pkg/auth"
	"github.com/go-chi/chi"
)

// NewHandler returns the api routes handler
func NewHandler(authService auth.Service) http.Handler {
	r := chi.NewRouter()

	// logger := httplog.NewLogger("login", httplog.Options{})
	// r.Use(httplog.RequestLogger(logger))

	r.Post("/", Login(authService))

	return r
}
