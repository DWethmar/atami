package router

// import (
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/dwethmar/atami/pkg/service"
// 	"github.com/dwethmar/atami/pkg/testutil"
// 	"github.com/stretchr/testify/assert"
// )

// func TestMessageAuthenticated(t *testing.T) {
// 	authService := service.NewAuthServiceMemory()
// 	handler := NewMessageRouter(authService)
// 	req := httptest.NewRequest("GET", "/", nil)
// 	testutil.WithAuthorizationHeader(req, authService)
// 	rr := httptest.NewRecorder()
// 	handler.ServeHTTP(rr, req)
// 	assert.Equal(t, http.StatusOK, rr.Code, rr.Body.String())
// }

// func TestMessageUnauthenticated(t *testing.T) {
// 	handler := NewMessageRouter(service.NewAuthServiceMemory())
// 	req := httptest.NewRequest("GET", "/feed", nil)
// 	rr := httptest.NewRecorder()
// 	handler.ServeHTTP(rr, req)
// 	assert.Equal(t, http.StatusUnauthorized, rr.Code, rr.Body.String())
// }
