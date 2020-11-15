package handler

import (
	"net/http"
)

// ListMessages handler
func ListMessages() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
}
