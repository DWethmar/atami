package middleware

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/dwethmar/atami/pkg/auth"
	"github.com/dwethmar/atami/pkg/domain"
	"github.com/dwethmar/atami/pkg/domain/user"
	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/stretchr/testify/assert"
)

func TestAuthenticated(t *testing.T) {
	store := domain.NewInMemoryStore(memstore.NewStore())

	var accessToken string
	if user, err := store.User.Create(user.CreateUser{
		Username: "userx",
		Email:    "test@test.nl",
		Password: "Abcd1234!@#$",
	}); err == nil {
		accessToken, err = auth.CreateAccessToken(user.UID, user.Username, 4100760000)
		if !assert.NoError(t, err) {
			return
		}
	} else {
		assert.Fail(t, "could not create user :<")
		return
	}

	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, err := GetUser(r.Context())
		assert.NoError(t, err)

		if assert.NotNil(t, user) && user.ID != 0 {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	})

	middleware := Authenticated(store.User)
	handlerToTest := middleware(nextHandler)

	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	rr := httptest.NewRecorder()
	handlerToTest.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code, rr.Body.String())
}

func TestAuthenticatedContext(t *testing.T) {
	user := &user.User{
		ID:        1,
		UID:       "abc123",
		Username:  "test_user",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	req := httptest.NewRequest("POST", "/", nil)
	ctx := req.Context()
	ctx = WithUser(ctx, user)

	userStored, err := GetUser(ctx)
	if assert.NoError(t, err) {
		assert.True(t, user.Equal(*userStored))
	}
}
