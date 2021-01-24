package middleware

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dwethmar/atami/pkg/api/response"
	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
)

func TestRequireUID(t *testing.T) {
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		UID, err := UIDFromContext(r.Context())

		if err != nil {
			t.Error(err)
		}

		if UID != "6f2128ce" {
			t.Error("wrong id")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
	})

	router := chi.NewRouter()
	router.Get("/{uid}", RequireUID(nextHandler))

	req := httptest.NewRequest("GET", "/6f2128ce", nil)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code, "Status code should be equal")
}

func TestRequireMissingUID(t *testing.T) {
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	router := chi.NewRouter()
	router.Get("/", RequireUID(nextHandler))

	req := httptest.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Code, "Status code should be equal")

	expectedResponds := response.ErrorResponds{
		Error:   http.StatusText(http.StatusBadRequest),
		Message: "no UID found in URL",
	}

	// Check the response body is what we expect.
	expected, _ := json.Marshal(expectedResponds)
	assert.Equal(t, string(expected), rr.Body.String(), "handler returned unexpected body")
}
