package feed

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dwethmar/atami/pkg/api/testutil"
	"github.com/dwethmar/atami/pkg/service"
	"github.com/stretchr/testify/assert"
)

func TestAuthenticated(t *testing.T) {
	authService := service.NewAuthServiceMemory()
	handler := NewHandler(authService)
	req := httptest.NewRequest("GET", "/", nil)
	testutil.WithAuthorizationHeader(req, authService)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code, rr.Body.String())
}

func TestUnauthenticated(t *testing.T) {
	handler := NewHandler(service.NewAuthServiceMemory())
	req := httptest.NewRequest("GET", "/feed", nil)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusUnauthorized, rr.Code, rr.Body.String())
}
