package login

import (
	"net/http"

	"github.com/dwethmar/atami/pkg/auth"
)

// Login handles login requests
func Login(service auth.Service) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}
