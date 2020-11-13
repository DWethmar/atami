package thread

import (
	"net/http"

	"github.com/dwethmar/atami/pkg/api/middleware"
	"github.com/dwethmar/atami/pkg/auth"
	"github.com/go-chi/chi"
	"github.com/go-chi/httplog"
)

// NewHandler returns the thread routes handler
func NewHandler(service auth.Service) http.Handler {
	r := chi.NewRouter()

	logger := httplog.NewLogger("thread", httplog.Options{})
	r.Use(httplog.RequestLogger(logger))
	r.Use(middleware.Token)

	r.Get("/", Thread(service))

	return r
}
