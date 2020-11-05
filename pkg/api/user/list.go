package user

import (
	"fmt"
	"net/http"

	"github.com/dwethmar/atami/pkg/api/response"
	"github.com/dwethmar/atami/pkg/usecase/userusecase"
)

// ListUsers handler
func ListUsers(usecase *userusecase.Usecase) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if users, err := usecase.List(); err == nil {
			response.SendJSON(w, r, users, 200)
		} else {
			fmt.Printf("Error: %v \n", err)
			response.SendServerError(w, r)
		}
	})
}
