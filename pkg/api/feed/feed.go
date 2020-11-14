package feed

import (
	"fmt"
	"net/http"

	"github.com/dwethmar/atami/pkg/api/middleware"
)

// Feed WIP
func Feed() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, err := middleware.UserFromContext(r.Context())
		if err != nil {
			fmt.Fprintf(w, ":(")
		}
		fmt.Fprintf(w, fmt.Sprintf("Hello %v", user.Username))
	})
}
