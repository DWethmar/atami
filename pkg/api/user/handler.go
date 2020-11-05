package user

import (
	"net/http"

	"github.com/dwethmar/atami/pkg/auth"
	"github.com/go-chi/chi"
)

// NewHandler returns the api routes handler
func NewHandler(service auth.Service) http.Handler {
	r := chi.NewRouter()

	r.Get("/users", ListUsers(service))
	r.Post("/register", RegisterUser(service))

	return r
}
