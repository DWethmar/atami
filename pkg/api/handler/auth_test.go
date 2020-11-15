package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"

	"github.com/dwethmar/atami/pkg/api/response"
	"github.com/dwethmar/atami/pkg/auth"
	"github.com/dwethmar/atami/pkg/service"
	"github.com/stretchr/testify/assert"
)

var users = []*auth.CreateUser{
	{
		Username: "Test1",
		Email:    "test1@test.com",
		Password: "abcd123!@#A",
	},
	{
		Username: "Test2",
		Email:    "test2@test.com",
		Password: "abcd123!@#A",
	},
}

func TestList(t *testing.T) {
	authService := service.NewAuthServiceMemory()

	var expectedResponds = make([]*Responds, len(users))
	for i, user := range users {
		r, _ := authService.Register(*user)
		expectedResponds[i] = toRespond(r)
	}

	req, err := http.NewRequest("GET", "/", nil)
	assert.Nil(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ListUsers(authService))
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "Status code should be equal")
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"), "Content-Type code should be equal")

	// Check the response body is what we expect.
	expected, _ := json.Marshal(expectedResponds)
	assert.Equal(t, string(expected), rr.Body.String(), "handler returned unexpected body")
}

var newUser = NewUser{
	Username: "Username",
	Email:    "test@test.com",
	Password: "myL!ttleSecr3t",
}

var invalidUser1 = NewUser{
	Username: "",
	Email:    "test@test.com",
	Password: "myL!ttleSecr3t",
}

func TestRegisterUser(t *testing.T) {
	authService := service.NewAuthServiceMemory()
	handler := http.HandlerFunc(Register(authService))

	body, _ := json.Marshal(newUser)
	req := httptest.NewRequest("POST", "/", bytes.NewBuffer(body))
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))

	// Check the response body is what we expect.
	createdUser := Responds{}
	assert.NoError(t, json.Unmarshal(rr.Body.Bytes(), &createdUser))
	assert.Equal(t, "Username", createdUser.Username)
}

func TestRegisterInvalidUser(t *testing.T) {
	service := service.NewAuthServiceMemory()
	handler := http.HandlerFunc(Register(service))

	requests := []*NewUser{
		{
			Username: "",
			Email:    "test@test.com",
			Password: "myL!ttleSecr3t",
		},
		{
			Username: "qwerty",
			Email:    "x",
			Password: "myL!ttleSecr3t",
		},
		{},
	}

	expectedErrors := []*response.ErrorResponds{
		{
			Error:   "Bad Request",
			Message: "username is required",
		},
		{
			Error:   "Bad Request",
			Message: "email is invalid",
		},
		{
			Error:   "Bad Request",
			Message: "username is required. email is required. password is required",
		},
	}

	for i, r := range requests {
		e := expectedErrors[i]

		body, _ := json.Marshal(r)
		req := httptest.NewRequest("POST", "/", bytes.NewBuffer(body))
		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))

		// Check the response body is what we expect.
		errorResponds := response.ErrorResponds{}
		assert.NoError(t, json.Unmarshal(rr.Body.Bytes(), &errorResponds))

		// Check the response body is what we expect.
		expected, _ := json.Marshal(e)
		assert.Equal(t, string(expected), rr.Body.String())
	}

}

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
