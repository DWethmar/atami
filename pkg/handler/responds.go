package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// ErrorResponds error responds type
type ErrorResponds struct {
	Error string `json:"error"`
}

// SendJSON set json responds.
func sendJSON(w http.ResponseWriter, r *http.Request, v interface{}) {
	w.Header().Add("Content-Type", "application/json")
	b, err := json.Marshal(v)
	if err != nil {
		log.Print(fmt.Sprintf("Error while encoding JSON: %v", err))
		io.WriteString(w, `{"error": "Internal server error"}`)
	} else {
		io.WriteString(w, string(b))
	}
}
