package feed

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dwethmar/atami/pkg/api/testutil"
	"github.com/dwethmar/atami/pkg/service"
)

func TestAuthenticated(t *testing.T) {
	authService := service.NewAuthServiceMemory()
	handler := NewHandler(authService)
	req := httptest.NewRequest("GET", "/", nil)
	testutil.TestWithAuthorizationHeader(t, authService, req, handler, http.StatusOK)
}

func TestUnauthenticated(t *testing.T) {
	handler := NewHandler(service.NewAuthServiceMemory())
	req := httptest.NewRequest("GET", "/feed", nil)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	testutil.TestStatus(t, req, handler, http.StatusUnauthorized)
}
