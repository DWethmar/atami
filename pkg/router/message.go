package router

// import (
// 	"net/http"

// 	"github.com/dwethmar/atami/pkg/api/handler/channel"

// 	"github.com/dwethmar/atami/pkg/api/middleware"
// 	"github.com/dwethmar/atami/pkg/auth"
// 	"github.com/go-chi/chi"
// 	"github.com/go-chi/httplog"
// )

// // NewMessageRouter creates new message router
// func NewMessageRouter(authService auth.Service) http.Handler {
// 	r := chi.NewRouter()

// 	logger := httplog.NewLogger("message", httplog.Options{})
// 	r.Use(httplog.RequestLogger(logger))
// 	r.Use(middleware.Authenticated(authService))

// 	r.Get("/", channel.Index())

// 	return r
// }
