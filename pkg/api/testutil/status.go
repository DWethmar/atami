package testutil

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestStatus test if handles authorization
func TestStatus(t *testing.T, req *http.Request, handler http.Handler, expectedStatus int) bool {
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	return assert.Equal(t, expectedStatus, rr.Code, rr.Body.String())
}
