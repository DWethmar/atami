package testutil

import (
	"fmt"
	"net/http"

	"github.com/dwethmar/atami/pkg/api/token"
	"github.com/dwethmar/atami/pkg/auth"
)

// WithAuthorizationHeader test if handles authorization
func WithAuthorizationHeader(req *http.Request, authService auth.Service) error {
	user, err := authService.Register(auth.CreateUser{
		Username: "test",
		Email:    "test@test.nl",
		Password: "Test1234!@#$",
	})
	if err != nil {
		return err
	}

	accessToken, err := token.CreateToken(user.UID, user.Username, 4100760000)
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken.AccessToken))
	return nil
}
