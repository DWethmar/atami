package router

import (
	"net/http"

	"github.com/dwethmar/atami/pkg/api/handler"
	"github.com/dwethmar/atami/pkg/auth"
	"github.com/go-chi/chi"
)

// NewLoginRouter creates a new login router
func NewLoginRouter(authService auth.Service) http.Handler {
	r := chi.NewRouter()

	r.Post("/", handler.Login(authService))

	return r
}
