package router

import (
	"net/http"

	"github.com/dwethmar/atami/pkg/api/handler"
	"github.com/dwethmar/atami/pkg/auth"

	"github.com/go-chi/chi"
)

// NewAuthRouter returns the api routes handler
func NewAuthRouter(authService *auth.Service) http.Handler {
	r := chi.NewRouter()

	r.Post("/register", handler.Register(authService))
	r.Post("/login", handler.Login(authService))

	return r
}
