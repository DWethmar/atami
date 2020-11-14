package router

import (
	"net/http"

	"github.com/dwethmar/atami/pkg/auth"
	"github.com/dwethmar/atami/pkg/message"
	"github.com/dwethmar/atami/pkg/message/handler"
	"github.com/dwethmar/atami/pkg/middleware"
	"github.com/go-chi/chi"
	"github.com/go-chi/httplog"
)

// NewMessageRouter creates new message router
func NewMessageRouter(authService auth.Service, messageService message.Service) http.Handler {
	r := chi.NewRouter()

	logger := httplog.NewLogger("message", httplog.Options{})
	r.Use(httplog.RequestLogger(logger))
	r.Use(middleware.Authenticated(authService))

	r.Get("/", handler.Index())

	return r
}
