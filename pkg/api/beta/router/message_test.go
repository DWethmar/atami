package router

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dwethmar/atami/pkg/service"
	"github.com/stretchr/testify/assert"
)

func TestMessageAuthenticated(t *testing.T) {
	userService, store := service.NewUserServiceMemory()
	authService := service.NewAuthServiceMemory(store)

	messageService := service.NewMessageServiceMemory()
	handler := NewMessageRouter(userService, messageService)

	req := httptest.NewRequest("GET", "/", nil)

	if err := WithAuthorizationHeader(req, authService); assert.NoError(t, err) {
		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusOK, rr.Code, rr.Body.String())
	}
}

func TestMessageUnauthenticated(t *testing.T) {
	userService, _ := service.NewUserServiceMemory()
	messageService := service.NewMessageServiceMemory()
	handler := NewMessageRouter(userService, messageService)

	req := httptest.NewRequest("GET", "/feed", nil)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusUnauthorized, rr.Code, rr.Body.String())
}
