package router

import (
	"net/http"

	"github.com/dwethmar/atami/pkg/api/handler"
	"github.com/dwethmar/atami/pkg/api/middleware"
	"github.com/dwethmar/atami/pkg/message"
	"github.com/dwethmar/atami/pkg/user"
	"github.com/go-chi/chi"
	"github.com/go-chi/httplog"
)

// NewMessageRouter creates new message router
func NewMessageRouter(userService *user.Service, messageService *message.Service) http.Handler {
	r := chi.NewRouter()

	logger := httplog.NewLogger("message", httplog.Options{})
	r.Use(httplog.RequestLogger(logger))
	r.Use(middleware.Authenticated(userService))

	r.Get("/", handler.ListMessages(messageService))
	r.Post("/", handler.CreateMessages(messageService))

	return r
}
