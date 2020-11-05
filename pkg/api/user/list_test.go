package user

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/dwethmar/atami/pkg/auth"
	userMem "github.com/dwethmar/atami/pkg/auth/memory"
	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/stretchr/testify/assert"
)

var now = time.Now()

var users = []*auth.User{
	{
		ID:        1,
		UID:       "1",
		Username:  "Test1",
		Email:     "test1@test.com",
		CreatedAt: now,
		UpdatedAt: now,
		Password:  "",
	},
	{
		ID:        2,
		UID:       "2",
		Username:  "Test2",
		Email:     "test2@test.com",
		CreatedAt: now,
		UpdatedAt: now,
		Password:  "",
	},
}

var expectedUsers = []*Responds{
	{
		UID:       "1",
		Username:  "Test1",
		CreatedAt: now,
	},
	{
		UID:       "2",
		Username:  "Test2",
		CreatedAt: now,
	},
}

func TestList(t *testing.T) {
	store := memstore.New()
	for _, user := range users {
		store.Add(user.ID.String(), *user)
	}

	service := userMem.NewService(store)

	req, err := http.NewRequest("GET", "/", nil)
	assert.Nil(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ListUsers(service))
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "Status code should be equal")
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"), "Content-Type code should be equal")

	// Check the response body is what we expect.
	expected, _ := json.Marshal(expectedUsers)
	assert.Equal(t, string(expected), rr.Body.String(), "handler returned unexpected body")
}
