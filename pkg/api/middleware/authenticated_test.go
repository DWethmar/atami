package middleware

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/dwethmar/atami/pkg/auth"
	"github.com/dwethmar/atami/pkg/service"
	"github.com/stretchr/testify/assert"
)

func TestAuthenticated(t *testing.T) {
	service := service.NewAuthServiceMemory()
	user, err := service.Register(auth.CreateUser{
		Username: "userx",
		Email:    "test@test.nl",
		Password: "Abcd1234!@#$",
	})
	if err != nil {
		assert.Fail(t, "could not create user :<")
		return
	}

	token, err := auth.CreateToken(user.UID, user.Username, 4100760000)
	if !assert.NoError(t, err) {
		return
	}

	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, err := UserFromContext(r.Context())
		assert.NoError(t, err)

		if assert.NotNil(t, user) && user.ID != 0 {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	})

	middleware := Authenticated(service)
	handlerToTest := middleware(nextHandler)

	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))

	rr := httptest.NewRecorder()
	handlerToTest.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code, rr.Body.String())
}

func TestAuthenticatedContext(t *testing.T) {
	user := &auth.User{
		ID:        1,
		UID:       "abc123",
		Username:  "test_user",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	req := httptest.NewRequest("POST", "/", nil)
	ctx := req.Context()
	ctx = WithUser(ctx, user)

	userStored, err := UserFromContext(ctx)
	if assert.NoError(t, err) {
		assert.True(t, user.Equal(*userStored))
	}
}
