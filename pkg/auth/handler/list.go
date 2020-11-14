package handler

import (
	"fmt"
	"net/http"

	"github.com/dwethmar/atami/pkg/auth"
	"github.com/dwethmar/atami/pkg/response"
)

// ListUsers handler
func ListUsers(service auth.Service) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if users, err := service.Find(); err == nil {
			response.SendJSON(w, r, toResponds(users), 200)
		} else {
			fmt.Printf("Error: %v \n", err)
			response.SendServerError(w, r)
		}
	})
}
