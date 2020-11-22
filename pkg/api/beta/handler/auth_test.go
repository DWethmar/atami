package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/dwethmar/atami/pkg/api/response"
	"github.com/dwethmar/atami/pkg/auth"
	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/dwethmar/atami/pkg/service"
	"github.com/stretchr/testify/assert"
)

// NewUser struct definition
type NewUser struct {
	Username string
	Email    string
	Password string
}

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
	userService, store := service.NewUserServiceMemory()
	authService := service.NewAuthServiceMemory(store)

	var expectedResponds = make([]*Responds, len(users))
	for i, user := range users {
		r, _ := authService.Register(*user)
		expectedResponds[i] = toRespond(r)
	}

	req, err := http.NewRequest("GET", "/", nil)
	assert.Nil(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ListUsers(userService))
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
	store := memstore.New()
	authService := service.NewAuthServiceMemory(store)
	handler := http.HandlerFunc(Register(authService))

	form := url.Values{}
	form.Add("email", newUser.Email)
	form.Add("password", newUser.Password)
	form.Add("username", newUser.Username)

	req := httptest.NewRequest("POST", "/", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code, rr.Body.String())
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))

	// Check the response body is what we expect.
	createdUser := Responds{}
	assert.NoError(t, json.Unmarshal(rr.Body.Bytes(), &createdUser))
	assert.Equal(t, "Username", createdUser.Username)
}

func TestRegisterInvalidUser(t *testing.T) {
	store := memstore.New()
	authService := service.NewAuthServiceMemory(store)
	handler := http.HandlerFunc(Register(authService))

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

		form := url.Values{}
		form.Add("email", r.Email)
		form.Add("password", r.Password)
		form.Add("username", r.Username)

		req := httptest.NewRequest("POST", "/", strings.NewReader(form.Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code, rr.Body.String())
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
	store := memstore.New()
	userService, store := service.NewUserServiceMemory()
	authService := service.NewAuthServiceMemory(store)

	_, err := authService.Register(auth.CreateUser{
		Username: "test_username",
		Email:    "test@test.com",
		Password: "test123!@#ABC",
	})
	assert.NoError(t, err)
	handler := http.HandlerFunc(Login(authService, userService))

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

func TestRefresh(t *testing.T) {
	os.Setenv("ACCESS_SECRET", "abc")
	store := memstore.New()
	userService, store := service.NewUserServiceMemory()
	authService := service.NewAuthServiceMemory(store)

	user, err := authService.Register(auth.CreateUser{
		Username: "test_username",
		Email:    "test@test.com",
		Password: "test123!@#ABC",
	})
	assert.NoError(t, err)

	handler := http.HandlerFunc(Refresh(authService, userService))

	refreshToken, err := auth.CreateRefreshToken(
		user.UID,
		"abcdefg",
		time.Now().Add(time.Hour*1).Unix(),
	)
	assert.NoError(t, err)

	req := httptest.NewRequest("POST", "/beta/auth/refresh", nil)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.AddCookie(createRefreshCookie(
		time.Now().Add(time.Hour*2),
		"localhost",
		refreshToken,
	))

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if assert.Equal(t, http.StatusOK, rr.Code, rr.Body.String()) {
		assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))
		// Check the response body is what we expect.
		responds := loginResponds{}
		assert.NoError(t, json.Unmarshal(rr.Body.Bytes(), &responds))
		assert.NotEmpty(t, responds.AccessToken)

		token, err := auth.VerifyAccessToken(responds.AccessToken)
		assert.NoError(t, err)
		assert.NotEmpty(t, token)

		claims, ok := token.Claims.(*auth.CustomClaims)
		if assert.True(t, ok) {
			assert.Equal(t, "abcdefg", claims.SessionID)
		}
	}
}

func TestInvalidRefresh(t *testing.T) {
	os.Setenv("ACCESS_SECRET", "abc")
	store := memstore.New()
	userService, store := service.NewUserServiceMemory()
	authService := service.NewAuthServiceMemory(store)

	handler := http.HandlerFunc(Refresh(authService, userService))

	req := httptest.NewRequest("POST", "/beta/auth/refresh", nil)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Code, rr.Body.String())

	responds := response.ErrorResponds{}
	assert.NoError(t, json.Unmarshal(rr.Body.Bytes(), &responds))
	assert.Equal(t, "Bad Request", responds.Error)
	assert.Equal(t, "", responds.Message)
}

func TestRefreshWithExpiredToken(t *testing.T) {
	os.Setenv("ACCESS_SECRET", "abc")
	store := memstore.New()
	userService, store := service.NewUserServiceMemory()
	authService := service.NewAuthServiceMemory(store)

	user, err := authService.Register(auth.CreateUser{
		Username: "test_username",
		Email:    "test@test.com",
		Password: "test123!@#ABC",
	})
	assert.NoError(t, err)

	handler := http.HandlerFunc(Refresh(authService, userService))

	refreshToken, err := auth.CreateRefreshToken(
		user.UID,
		"abcdefg",
		667184461,
	)
	assert.NoError(t, err)

	req := httptest.NewRequest("POST", "/beta/auth/refresh", nil)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.AddCookie(createRefreshCookie(
		time.Now().Add(time.Hour*2),
		refreshToken,
	))

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if assert.Equal(t, http.StatusBadRequest, rr.Code, rr.Body.String()) {
		assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))
		// Check the response body is what we expect.
		responds := response.ErrorResponds{}
		assert.NoError(t, json.Unmarshal(rr.Body.Bytes(), &responds))
		assert.Equal(t, "Bad Request", responds.Error)
		assert.Equal(t, "token is expired", responds.Message)
	}
}