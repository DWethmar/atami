package registration

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dwethmar/atami/pkg/auth"
	userMem "github.com/dwethmar/atami/pkg/auth/memory"
	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/stretchr/testify/assert"
)

var users = []*auth.RegisterUser{
	{
		Username:      "Test1",
		Email:         "test1@test.com",
		PlainPassword: "abcd123!@#A",
	},
	{
		Username:      "Test2",
		Email:         "test2@test.com",
		PlainPassword: "abcd123!@#A",
	},
}

func TestList(t *testing.T) {
	service := userMem.NewService(memstore.New())

	var expectedResponds = make([]*Responds, len(users))
	for i, user := range users {
		r, _ := service.Register(*user)
		expectedResponds[i] = toRespond(r)
	}

	req, err := http.NewRequest("GET", "/", nil)
	assert.Nil(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ListUsers(service))
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "Status code should be equal")
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"), "Content-Type code should be equal")

	// Check the response body is what we expect.
	expected, _ := json.Marshal(expectedResponds)
	assert.Equal(t, string(expected), rr.Body.String(), "handler returned unexpected body")
}
