package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

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
