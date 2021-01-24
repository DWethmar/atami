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

// JSON set json responds.
func JSON(w http.ResponseWriter, r *http.Request, v interface{}, code int) {
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

// ServerError set server error
func ServerError(w http.ResponseWriter, r *http.Request) {
	JSON(w, r, ErrorResponds{
		Error: http.StatusText(http.StatusInternalServerError),
	}, http.StatusInternalServerError)
}

// BadRequestError set bad request responds
func BadRequestError(w http.ResponseWriter, r *http.Request, err error) {
	JSON(w, r, ErrorResponds{
		Error:   http.StatusText(http.StatusBadRequest),
		Message: err.Error(),
	}, http.StatusBadRequest)
}

// UnauthorizedError set bad request responds
func UnauthorizedError(w http.ResponseWriter, r *http.Request, err error) {
	JSON(w, r, ErrorResponds{
		Error:   http.StatusText(http.StatusUnauthorized),
		Message: err.Error(),
	}, http.StatusUnauthorized)
}

// NotFoundError set not found responds
func NotFoundError(w http.ResponseWriter, r *http.Request) {
	JSON(w, r, ErrorResponds{
		Error: "Resource not found.",
	}, http.StatusNotFound)
}
