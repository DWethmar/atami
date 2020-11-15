package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/dwethmar/atami/pkg/api/middleware"
	"github.com/dwethmar/atami/pkg/api/response"
)

// ListMessages handler
func ListMessages() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		usr, err := middleware.UserFromContext(r.Context())
		if err != nil || usr == nil {
			response.SendBadRequestError(w, r, errors.New(":P"))
			return
		}

		fmt.Fprintf(w, "Hi there %v, I love %s!", usr.Username, r.URL.Path[1:])
	})
}
