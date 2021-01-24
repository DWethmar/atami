package beta

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/dwethmar/atami/pkg/api/middleware"
	"github.com/dwethmar/atami/pkg/api/response"
	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/dwethmar/atami/pkg/message"
	"github.com/dwethmar/atami/pkg/user/memory/util"

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
				ID:       1,
				UID:      "UID1",
				Username: "test",
			},
		},
		{
			ID:              2,
			UID:             "abcdefg12345",
			Text:            "lorum ipsum 2",
			CreatedAt:       time.Now(),
			CreatedByUserID: 1,
			User: &message.User{
				ID:       1,
				UID:      "UID1",
				Username: "test",
			},
		},
	}
	store := memstore.NewStore()
	user := util.AddTestUser(store, 1)

	for _, msg := range messages {
		store.GetMessages().Put(strconv.Itoa(msg.ID), *msg)
	}
	ms := service.NewMessageServiceMemory(store)
	req, err := http.NewRequest("GET", "/", nil)
	assert.NoError(t, err)

	// Add user to context
	ctx := req.Context()
	ctx = middleware.WithUser(ctx, user)
	req = req.WithContext(ctx)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ListMessages(ms))
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))

	expectedResponds := make([]*Message, len(messages))
	for i, m := range messages {
		expectedResponds[i] = mapMessage(m)
	}

	// Check the response body is what we expect.
	expected, _ := json.Marshal(expectedResponds)
	assert.Equal(t, string(expected), rr.Body.String(), "handler returned unexpected body")
}

func TestGetMessage(t *testing.T) {
	messages := []*message.Message{
		{
			ID:              1,
			UID:             "abcdefg1234",
			Text:            "lorum ipsum",
			CreatedAt:       time.Now(),
			CreatedByUserID: 1,
			User: &message.User{
				ID:       1,
				UID:      "UID1",
				Username: "test",
			},
		},
		{
			ID:              2,
			UID:             "abcdefg12345",
			Text:            "lorum ipsum 2",
			CreatedAt:       time.Now(),
			CreatedByUserID: 1,
			User: &message.User{
				ID:       1,
				UID:      "UID1",
				Username: "test",
			},
		},
	}
	store := memstore.NewStore()
	user := util.AddTestUser(store, 1)

	for _, msg := range messages {
		store.GetMessages().Put(strconv.Itoa(msg.ID), *msg)
	}
	ms := service.NewMessageServiceMemory(store)
	req, err := http.NewRequest("GET", "/abcdefg1234", nil)
	assert.NoError(t, err)

	// Add user to context
	ctx := req.Context()
	ctx = middleware.WithUser(ctx, user)
	req = req.WithContext(ctx)

	// Add UID to context
	ctx = req.Context()
	ctx = middleware.WithUID(ctx, "abcdefg1234")
	req = req.WithContext(ctx)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetMessage(ms))
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))

	expectedResponds := mapMessage(messages[0])

	// Check the response body is what we expect.
	expected, _ := json.Marshal(expectedResponds)
	assert.Equal(t, string(expected), rr.Body.String(), "handler returned unexpected body")
}

func TestNotFoundGetMessage(t *testing.T) {
	store := memstore.NewStore()
	user := util.AddTestUser(store, 1)

	ms := service.NewMessageServiceMemory(store)
	req, err := http.NewRequest("GET", "/notexistinguid", nil)
	assert.NoError(t, err)

	// Add user to context
	ctx := req.Context()
	ctx = middleware.WithUser(ctx, user)
	req = req.WithContext(ctx)

	// Add UID to context
	ctx = req.Context()
	ctx = middleware.WithUID(ctx, "notexistinguid")
	req = req.WithContext(ctx)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetMessage(ms))
	handler.ServeHTTP(rr, req)

	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))
	assert.Equal(t, http.StatusNotFound, rr.Code)

	expectedResponds := response.ErrorResponds{
		Error:   http.StatusText(http.StatusNotFound),
		Message: "",
	}
	// Check the response body is what we expect.
	expected, _ := json.Marshal(expectedResponds)
	assert.Equal(t, string(expected), rr.Body.String(), "handler returned unexpected body")
}

func TestCreateMessage(t *testing.T) {
	store := memstore.NewStore()
	user := util.AddTestUser(store, 1)
	ms := service.NewMessageServiceMemory(store)

	addEntry := CreatMessageInput{
		Text: "lorum ipsum",
	}
	body, _ := json.Marshal(addEntry)
	req := httptest.NewRequest("POST", "/", bytes.NewBuffer(body))

	// Add user to context.
	ctx := req.Context()
	ctx = middleware.WithUser(ctx, user)
	req = req.WithContext(ctx)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateMessage(ms))
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))

	// Check the response body is what we expect.
	result := CreatMessageSuccess{}
	assert.Nil(t, json.Unmarshal(rr.Body.Bytes(), &result))
}

func TestUnauthorizedCreateMessage(t *testing.T) {
	store := memstore.NewStore()
	ms := service.NewMessageServiceMemory(store)

	addEntry := CreatMessageInput{
		Text: "sadsdkjskjdskjsjdsjkskjkjdkjkjsdkjjdsk",
	}
	body, _ := json.Marshal(addEntry)
	req := httptest.NewRequest("POST", "/", bytes.NewBuffer(body))

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateMessage(ms))
	handler.ServeHTTP(rr, req)

	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))
	assert.Equal(t, http.StatusUnauthorized, rr.Code)

	expectedResponds := response.ErrorResponds{
		Error:   http.StatusText(http.StatusUnauthorized),
		Message: "unauthorized",
	}
	// Check the response body is what we expect.
	expected, _ := json.Marshal(expectedResponds)
	assert.Equal(t, string(expected), rr.Body.String(), "handler returned unexpected body")
}

func TestDeleteMessage(t *testing.T) {
	messages := []*message.Message{
		{
			ID:              1,
			UID:             "abcdefg1234",
			Text:            "lorum ipsum",
			CreatedAt:       time.Now(),
			CreatedByUserID: 1,
			User: &message.User{
				ID:       1,
				UID:      "UID1",
				Username: "test",
			},
		},
	}
	store := memstore.NewStore()
	user := util.AddTestUser(store, 1)

	for _, msg := range messages {
		store.GetMessages().Put(strconv.Itoa(msg.ID), *msg)
	}
	ms := service.NewMessageServiceMemory(store)
	req, err := http.NewRequest("DELETE", "/abcdefg1234", nil)
	assert.NoError(t, err)

	// Add user to context
	ctx := req.Context()
	ctx = middleware.WithUser(ctx, user)
	req = req.WithContext(ctx)

	// Add UID to context
	ctx = req.Context()
	ctx = middleware.WithUID(ctx, "abcdefg1234")
	req = req.WithContext(ctx)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetMessage(ms))
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))

	expectedResponds := mapMessage(messages[0])

	// Check the response body is what we expect.
	expected, _ := json.Marshal(expectedResponds)
	assert.Equal(t, string(expected), rr.Body.String(), "handler returned unexpected body")
}

func TestUnauthorizedDeleteMessage(t *testing.T) {
	messages := []*message.Message{
		{
			ID:              1,
			UID:             "abcdefg1234",
			Text:            "lorum ipsum",
			CreatedAt:       time.Now(),
			CreatedByUserID: 2,
			User: &message.User{
				ID:       2,
				UID:      "UID1",
				Username: "test",
			},
		},
	}
	store := memstore.NewStore()
	user := util.AddTestUser(store, 1)
	util.AddTestUser(store, 2)

	for _, msg := range messages {
		store.GetMessages().Put(strconv.Itoa(msg.ID), *msg)
	}
	ms := service.NewMessageServiceMemory(store)
	req, err := http.NewRequest("DELETE", "/abcdefg1234", nil)
	assert.NoError(t, err)

	// Add user to context
	ctx := req.Context()
	ctx = middleware.WithUser(ctx, user)
	req = req.WithContext(ctx)

	// Add UID to context
	ctx = req.Context()
	ctx = middleware.WithUID(ctx, "abcdefg1234")
	req = req.WithContext(ctx)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(DeleteMessage(ms))
	handler.ServeHTTP(rr, req)

	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))
	assert.Equal(t, http.StatusUnauthorized, rr.Code)

	expectedResponds := response.ErrorResponds{
		Error:   http.StatusText(http.StatusUnauthorized),
		Message: "unauthorized",
	}
	// Check the response body is what we expect.
	expected, _ := json.Marshal(expectedResponds)
	assert.Equal(t, string(expected), rr.Body.String(), rr.Body.String())
}

func TestNotFoundDeleteMessage(t *testing.T) {
	store := memstore.NewStore()
	user := util.AddTestUser(store, 1)

	ms := service.NewMessageServiceMemory(store)
	req, err := http.NewRequest("DELETE", "/abcdefg1234", nil)
	assert.NoError(t, err)

	// Add user to context
	ctx := req.Context()
	ctx = middleware.WithUser(ctx, user)
	req = req.WithContext(ctx)

	// Add UID to context
	ctx = req.Context()
	ctx = middleware.WithUID(ctx, "abcdefg1234")
	req = req.WithContext(ctx)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(DeleteMessage(ms))
	handler.ServeHTTP(rr, req)

	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))
	assert.Equal(t, http.StatusNotFound, rr.Code)
}
