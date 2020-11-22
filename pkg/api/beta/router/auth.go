package router

import (
	"net/http"

	"github.com/dwethmar/atami/pkg/api/beta/handler"
	"github.com/dwethmar/atami/pkg/auth"
	"github.com/dwethmar/atami/pkg/user"

	"github.com/go-chi/chi"
	"github.com/go-chi/httplog"
)

// NewAuthRouter returns the api routes handler
func NewAuthRouter(authService *auth.Service, userService *user.Service) http.Handler {
	r := chi.NewRouter()

	logger := httplog.NewLogger("auth", httplog.Options{})
	r.Use(httplog.RequestLogger(logger))

	r.Post("/register", handler.Register(authService))
	r.Post("/login", handler.Login(authService, userService))
	r.Post("/refresh", handler.Refresh(authService, userService))

	return r
}
