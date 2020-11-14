package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"

	"github.com/dwethmar/atami/pkg/auth"
	"github.com/dwethmar/atami/pkg/service"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	os.Setenv("ACCESS_SECRET", "abc")

	service := service.NewAuthServiceMemory()
	_, err := service.Register(auth.CreateUser{
		Username: "test_username",
		Email:    "test@test.com",
		Password: "test123!@#ABC",
	})
	assert.NoError(t, err)
	handler := http.HandlerFunc(Login(service))

	form := url.Values{}
	form.Add("email", "test@test.com")
	form.Add("password", "test123!@#ABC")

	req := httptest.NewRequest("POST", "/auth/login", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if assert.Equal(t, http.StatusOK, rr.Code, rr.Body.String()) {
		assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))
		// Check the response body is what we expect.
		responds := loginResponds{}
		assert.NoError(t, json.Unmarshal(rr.Body.Bytes(), &responds))
		assert.NotEmpty(t, responds.AccessToken)
	}
}
