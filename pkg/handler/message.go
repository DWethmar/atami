package handler

import "net/http"

// ListMessages creates a message list handler.
func ListMessages() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.WriteHeader(http.StatusOK)

		sendJSON(w, r, struct {
			Lala string `json:"anus"`
		}{
			Lala: "asd",
		})
	})
}
