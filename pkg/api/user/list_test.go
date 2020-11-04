package user

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/dwethmar/atami/pkg/model"
	"github.com/dwethmar/atami/pkg/usecase/userusecase"

	"github.com/stretchr/testify/assert"

	userMem "github.com/dwethmar/atami/pkg/user/memory"
)

var users = []*model.User{
	{
		ID:        1,
		UID:       "1",
		Username:  "Test1",
		CreatedAt: time.Now(),
	},
	{
		ID:        2,
		UID:       "2",
		Username:  "Test2",
		CreatedAt: time.Now(),
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
		store.Add(strconv.FormatInt(user.ID, 10), *user)
	}

	service := userMem.NewService(store)
	userUsecase := userusecase.NewUserUsecase(service)

	req, err := http.NewRequest("GET", "/", nil)
	assert.Nil(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ListUsers(userUsecase))
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "Status code should be equal")
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"), "Content-Type code should be equal")

	// Check the response body is what we expect.
	expected, _ := json.Marshal(toUsers(users))
	assert.Equal(t, string(expected), rr.Body.String(), "handler returned unexpected body")
}
