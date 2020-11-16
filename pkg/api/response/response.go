package response

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// ErrorResponds error responds type
type ErrorResponds struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

// SendJSON set json responds.
func SendJSON(w http.ResponseWriter, r *http.Request, v interface{}, code int) {
	w.Header().Set("Content-Type", "application/json")

	if b, err := json.Marshal(v); err != nil {
		log.Print(fmt.Sprintf("Error while encoding JSON: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"error": "Internal server error"}`)
	} else {
		w.WriteHeader(code)
		w.Write(b)
	}
}

// SendServerError set server error
func SendServerError(w http.ResponseWriter, r *http.Request) {
	SendJSON(w, r, ErrorResponds{
		Error: http.StatusText(http.StatusInternalServerError),
	}, http.StatusInternalServerError)
}

// SendBadRequestError set bad request responds
func SendBadRequestError(w http.ResponseWriter, r *http.Request, err error) {
	SendJSON(w, r, ErrorResponds{
		Error:   http.StatusText(http.StatusBadRequest),
		Message: err.Error(),
	}, http.StatusBadRequest)
}

// SendUnauthorizedError set bad request responds
func SendUnauthorizedError(w http.ResponseWriter, r *http.Request, err error) {
	SendJSON(w, r, ErrorResponds{
		Error:   http.StatusText(http.StatusUnauthorized),
		Message: err.Error(),
	}, http.StatusUnauthorized)
}

// SendNotFoundError set not found responds
func SendNotFoundError(w http.ResponseWriter, r *http.Request) {
	SendJSON(w, r, ErrorResponds{
		Error: "Resource not found.",
	}, http.StatusNotFound)
}
