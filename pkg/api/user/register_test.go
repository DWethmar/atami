package user

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dwethmar/atami/pkg/api/response"
	userMemory "github.com/dwethmar/atami/pkg/auth/memory"
	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/stretchr/testify/assert"
)

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
	service := userMemory.NewService(memstore.New())
	handler := http.HandlerFunc(RegisterUser(service))

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
	service := userMemory.NewService(memstore.New())
	handler := http.HandlerFunc(RegisterUser(service))

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
