package user

import (
	"net/http"

	"github.com/dwethmar/atami/pkg/usecase/userusecase"
	"github.com/go-chi/chi"
)

// NewHandler returns the api routes handler
func NewHandler(userUsecase *userusecase.Usecase) http.Handler {
	r := chi.NewRouter()

	r.Get("/users", ListUsers(userUsecase))
	r.Post("/register", RegisterUser(userUsecase))

	return r
}
