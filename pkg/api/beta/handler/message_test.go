package handler

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/dwethmar/atami/pkg/message"
	"github.com/dwethmar/atami/pkg/service"
	"github.com/stretchr/testify/assert"
)

// ListMessages handler
func TestListMessages(t *testing.T) {
	messages := []*message.Message{
		{
			ID:              1,
			UID:             "abcdefg1234",
			Text:            "lorum ipsum",
			CreatedAt:       time.Now(),
			CreatedByUserID: 1,
			User: &message.User{
				ID: 1,
			},
		},
		{
			ID:              2,
			UID:             "abcdefg12345",
			Text:            "lorum ipsum 2",
			CreatedAt:       time.Now(),
			CreatedByUserID: 1,
			User: &message.User{
				ID: 1,
			},
		},
	}
	store := memstore.New()
	for _, msg := range messages {
		store.Put(strconv.Itoa(msg.ID), msg)
	}
	ms := service.NewMessageServiceMemory(store)

	req, err := http.NewRequest("GET", "/", nil)
	assert.Nil(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ListMessages(ms))
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "Status code should be equal")
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"), "Content-Type code should be equal")

	// // Check the response body is what we expect.
	// expected, _ := json.Marshal(expectedResponds)
	// assert.Equal(t, string(expected), rr.Body.String(), "handler returned unexpected body")
}

func TestCreateMessage(t *testing.T) {

}
