package user

import (
	"net/http"

	"github.com/dwethmar/atami/pkg/user"
	"github.com/go-chi/chi"
)

// NewHandler returns the api routes handler
func NewHandler(finder *user.Finder) http.Handler {
	r := chi.NewRouter()

	r.Get("/", ListUsers(finder))

	return r
}
