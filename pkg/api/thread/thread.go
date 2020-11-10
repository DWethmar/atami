package thread

import (
	"fmt"
	"net/http"

	"github.com/dwethmar/atami/pkg/auth"
)

// Thread WIP
func Thread(service auth.Service) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})
}
