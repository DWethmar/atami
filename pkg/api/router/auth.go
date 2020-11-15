package router

import (
	"net/http"

	"github.com/dwethmar/atami/pkg/api/handler"
	"github.com/dwethmar/atami/pkg/auth"
	"github.com/dwethmar/atami/pkg/user"

	"github.com/go-chi/chi"
)

// NewAuthRouter returns the api routes handler
func NewAuthRouter(authService *auth.Service, userService *user.Service) http.Handler {
	r := chi.NewRouter()

	r.Post("/register", handler.Register(authService))
	r.Post("/login", handler.Login(authService, userService))

	return r
}
