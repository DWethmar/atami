package beta

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dwethmar/atami/pkg/auth"
	"github.com/stretchr/testify/assert"
)

// WithAuthorizationHeader test if handles authorization
func WithAuthorizationHeader(req *http.Request, authService *auth.Service) error {
	user, err := authService.Register(auth.CreateUser{
		Username: "test",
		Email:    "test@test.nl",
		Password: "Test1234!@#$",
	})
	if err != nil {
		return err
	}

	accessToken, err := auth.CreateAccessToken(user.UID, user.Username, 4100760000)
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	return nil
}

// TestStatus tests if status is returned by handler
func TestStatus(t *testing.T, req *http.Request, handler http.Handler, expectedStatus int) bool {
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	return assert.Equal(t, expectedStatus, rr.Code, rr.Body.String())
}
