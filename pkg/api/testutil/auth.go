package testutil

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/dwethmar/atami/pkg/api/token"
	"github.com/dwethmar/atami/pkg/auth"
	"github.com/stretchr/testify/assert"
)

// TestWithAuthorizationHeader test if handles authorization
func TestWithAuthorizationHeader(t *testing.T, authService auth.Service, req *http.Request, handler http.Handler, expectedStatus int) bool {
	user, err := authService.Register(auth.CreateUser{
		Username: "test",
		Email:    "test@test.nl",
		Password: "Test1234!@#$",
	})
	if !assert.NoError(t, err) {
		return false
	}

	accessToken, err := token.CreateToken(user.UID, user.Username, 4100760000)
	if !assert.NoError(t, err) {
		return false
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken.AccessToken))

	return TestStatus(t, req, handler, expectedStatus)
}
