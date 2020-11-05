package user

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	userMemory "github.com/dwethmar/atami/pkg/auth/memory"
	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/stretchr/testify/assert"
)

var newUser = NewUser{
	Username: "Username",
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

	assert.Equal(t, http.StatusCreated, rr.Code, "Status code should be equal")
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"), "Content-Type code should be equal")

	// Check the response body is what we expect.
	addedEntry := User{}
	assert.NoError(t, json.Unmarshal(rr.Body.Bytes(), &addedEntry))
	assert.Equal(t, "Username", addedEntry.Username)
}
