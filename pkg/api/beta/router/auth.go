package router

import (
	"net/http"

	"github.com/dwethmar/atami/pkg/api/beta/handler"
	"github.com/dwethmar/atami/pkg/auth"
	"github.com/dwethmar/atami/pkg/user"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/go-chi/httplog"
)

// NewAuthRouter returns the api routes handler
func NewAuthRouter(authService *auth.Service, userService *user.Service) http.Handler {
	r := chi.NewRouter()

	logger := httplog.NewLogger("auth", httplog.Options{})
	r.Use(httplog.RequestLogger(logger))
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
		// AllowedOrigins: []string{"*"},
		AllowOriginFunc: func(r *http.Request, origin string) bool {
			return true
		},
		AllowedMethods:   []string{"POST"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Post("/register", handler.Register(authService))
	r.Post("/login", handler.Login(authService, userService))
	r.Post("/refresh", handler.Refresh(authService, userService))

	return r
}
