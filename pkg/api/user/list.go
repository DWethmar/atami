package user

import (
	"fmt"
	"net/http"

	"github.com/dwethmar/atami/pkg/api/response"
	"github.com/dwethmar/atami/pkg/auth"
)

// ListUsers handler
func ListUsers(service auth.Service) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if users, err := service.FindAll(); err == nil {
			response.SendJSON(w, r, toResponds(users), 200)
		} else {
			fmt.Printf("Error: %v \n", err)
			response.SendServerError(w, r)
		}
	})
}
