package router

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dwethmar/atami/pkg/service"
	"github.com/stretchr/testify/assert"
)

func TestMessageAuthenticated(t *testing.T) {
	authService := service.NewAuthServiceMemory()
	messageService := service.NewMessageServiceMemory()
	handler := NewMessageRouter(authService, messageService)

	req := httptest.NewRequest("GET", "/", nil)
	WithAuthorizationHeader(req, authService)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code, rr.Body.String())
}

func TestMessageUnauthenticated(t *testing.T) {
	authService := service.NewAuthServiceMemory()
	messageService := service.NewMessageServiceMemory()
	handler := NewMessageRouter(authService, messageService)

	req := httptest.NewRequest("GET", "/feed", nil)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusUnauthorized, rr.Code, rr.Body.String())
}
