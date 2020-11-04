package user

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/dwethmar/atami/pkg/user"

	"github.com/stretchr/testify/assert"

	userMem "github.com/dwethmar/atami/pkg/user/memory"
)

var users = []*user.User{
	{
		ID:        1,
		UID:       "1",
		Username:  "Test1",
		Email:     "test1@test.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Password:  "",
	},
	{
		ID:        2,
		UID:       "2",
		Username:  "Test2",
		Email:     "test2@test.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Password:  "",
	},
}

func TestMapUsers(t *testing.T) {
	for i, user := range toUsers(users) {
		assert.Equal(t, users[i].UID, user.UID)
		assert.Equal(t, users[i].Username, user.Username)
	}
}

func TestList(t *testing.T) {
	store := memstore.New()
	for _, user := range users {
		store.Add(user.ID.String(), *user)
	}
	finder := userMem.NewFinder(store)

	req, err := http.NewRequest("GET", "/", nil)
	assert.Nil(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ListUsers(finder))
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "Status code should be equal")
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"), "Content-Type code should be equal")

	// Check the response body is what we expect.
	expected, _ := json.Marshal(toUsers(users))
	assert.Equal(t, string(expected), rr.Body.String(), "handler returned unexpected body")
}
