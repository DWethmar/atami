package registration

import (
	"net/http"

	"github.com/dwethmar/atami/pkg/auth"
	"github.com/go-chi/chi"
)

// NewHandler returns the api routes handler
func NewHandler(service auth.Service) http.Handler {
	r := chi.NewRouter()

	// logger := httplog.NewLogger("registration", httplog.Options{})
	// r.Use(httplog.RequestLogger(logger))

	r.Get("/users", ListUsers(service))
	r.Post("/", Register(service))

	return r
}
