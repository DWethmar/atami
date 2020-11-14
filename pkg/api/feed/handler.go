package feed

import (
	"net/http"

	"github.com/dwethmar/atami/pkg/api/middleware"
	"github.com/dwethmar/atami/pkg/auth"
	"github.com/go-chi/chi"
	"github.com/go-chi/httplog"
)

// NewHandler returns the thread routes handler
func NewHandler(authService auth.Service) http.Handler {
	r := chi.NewRouter()

	logger := httplog.NewLogger("feed", httplog.Options{})
	r.Use(httplog.RequestLogger(logger))
	r.Use(middleware.Authenticated(authService))

	r.Get("/", Feed())

	return r
}
