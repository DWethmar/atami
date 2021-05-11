package beta

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/dwethmar/atami/pkg/api/response"
	"github.com/dwethmar/atami/pkg/auth"
	"github.com/dwethmar/atami/pkg/domain"
	"github.com/dwethmar/atami/pkg/domain/message"
	"github.com/dwethmar/atami/pkg/domain/user"
	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/stretchr/testify/assert"
)

// ListMessages handler
func TestListMessages(t *testing.T) {
	store := domain.NewInMemoryStore(memstore.NewStore())
	user, err := store.User.Create(user.CreateUser{
		Username: "test",
		Email:    "test@test.nl",
		Password: "ABC123PXPXdddd@",
	})

	messages := []*message.Message{
		{
			ID:              1,
			UID:             "<to be replaced>",
			Text:            "lorum ipsum",
			CreatedAt:       time.Now(),
			CreatedByUserID: user.ID,
			User: &message.User{
				ID:       user.ID,
				UID:      user.UID,
				Username: user.Username,
			},
		},
		{
			ID:              2,
			UID:             "<to be replaced>",
			Text:            "lorum ipsum 2",
			CreatedAt:       time.Now(),
			CreatedByUserID: user.ID,
			User: &message.User{
				ID:       user.ID,
				UID:      user.UID,
				Username: user.Username,
			},
		},
	}

	for i, msg := range messages {
		m, err := store.Message.Create(message.CreateMessage{
			UID:             msg.UID,
			Text:            msg.Text,
			CreatedByUserID: msg.CreatedByUserID,
			CreatedAt:       msg.CreatedAt,
		})

		if assert.NoError(t, err) {
			messages[i].UID = m.UID
			messages[i].CreatedAt = m.CreatedAt
		} else {
			return
		}
	}

	req, err := http.NewRequest("GET", "/", nil)
	assert.NoError(t, err)

	token, _ := auth.CreateAccessToken(user.UID, "1", time.Now().Add(1000000).Unix())
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	rr := httptest.NewRecorder()
	NewMessageRouter(store).ServeHTTP(rr, req)

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
	store := domain.NewInMemoryStore(memstore.NewStore())
	user, err := store.User.Create(user.CreateUser{
		Username: "test",
		Email:    "test@test.nl",
		Password: "ABC123PXPXdddd@",
	})

	if err != nil {
		assert.Fail(t, err.Error())
	}

	messages := []*message.Message{
		{
			ID:              1,
			UID:             "<to be replaced>",
			Text:            "lorum ipsum",
			CreatedAt:       time.Now(),
			CreatedByUserID: user.ID,
			User: &message.User{
				ID:       user.ID,
				UID:      user.UID,
				Username: user.Username,
			},
		},
		{
			ID:              2,
			UID:             "<to be replaced>",
			Text:            "lorum ipsum 2",
			CreatedAt:       time.Now(),
			CreatedByUserID: user.ID,
			User: &message.User{
				ID:       user.ID,
				UID:      user.UID,
				Username: user.Username,
			},
		},
	}

	for i, msg := range messages {
		m, err := store.Message.Create(message.CreateMessage{
			UID:             msg.UID,
			Text:            msg.Text,
			CreatedByUserID: msg.CreatedByUserID,
			CreatedAt:       msg.CreatedAt,
		})

		if err != nil {
			assert.Fail(t, err.Error())
		}

		messages[i].UID = m.UID
		messages[i].CreatedAt = m.CreatedAt
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("/%s", messages[0].UID), nil)
	assert.NoError(t, err)

	token, _ := auth.CreateAccessToken(user.UID, "1", time.Now().Add(1000000).Unix())
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	rr := httptest.NewRecorder()
	NewMessageRouter(store).ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))

	expectedResponds := mapMessage(messages[0])

	// Check the response body is what we expect.
	expected, _ := json.Marshal(expectedResponds)
	assert.Equal(t, string(expected), rr.Body.String(), "handler returned unexpected body")
}

func TestNotFoundGetMessage(t *testing.T) {
	store := domain.NewInMemoryStore(memstore.NewStore())
	user, err := store.User.Create(user.CreateUser{
		Username: "test",
		Email:    "test@test.nl",
		Password: "ABC123PXPXdddd@",
	})

	req, err := http.NewRequest("GET", "/notexistinguid", nil)
	assert.NoError(t, err)

	token, _ := auth.CreateAccessToken(user.UID, "1", time.Now().Add(1000000).Unix())
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	rr := httptest.NewRecorder()
	NewMessageRouter(store).ServeHTTP(rr, req)

	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))
	assert.Equal(t, http.StatusNotFound, rr.Code)

	expectedResponds := response.ErrorResponds{
		Error:   http.StatusText(http.StatusNotFound),
		Message: "resource not found or unavailable",
	}
	// Check the response body is what we expect.
	expected, _ := json.Marshal(expectedResponds)
	assert.Equal(t, string(expected), rr.Body.String(), "handler returned unexpected body")
}

func TestCreateMessage(t *testing.T) {
	store := domain.NewInMemoryStore(memstore.NewStore())
	user, _ := store.User.Create(user.CreateUser{
		Username: "test",
		Email:    "test@test.nl",
		Password: "ABC123PXPXdddd@",
	})

	addEntry := CreatMessageInput{
		Text: "lorum ipsum",
	}

	body, _ := json.Marshal(addEntry)
	req := httptest.NewRequest("POST", "/", bytes.NewBuffer(body))

	token, _ := auth.CreateAccessToken(user.UID, "1", time.Now().Add(1000000).Unix())
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	rr := httptest.NewRecorder()
	NewMessageRouter(store).ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))

	// Check the response body is what we expect.
	result := CreatMessageSuccess{}
	assert.Nil(t, json.Unmarshal(rr.Body.Bytes(), &result))
}

func TestUnauthorizedCreateMessage(t *testing.T) {
	store := domain.NewInMemoryStore(memstore.NewStore())

	addEntry := CreatMessageInput{
		Text: "sadsdkjskjdskjsjdsjkskjkjdkjkjsdkjjdsk",
	}
	body, _ := json.Marshal(addEntry)
	req := httptest.NewRequest("POST", "/", bytes.NewBuffer(body))

	rr := httptest.NewRecorder()
	NewMessageRouter(store).ServeHTTP(rr, req)

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
	store := domain.NewInMemoryStore(memstore.NewStore())
	user, _ := store.User.Create(user.CreateUser{
		Username: "test",
		Email:    "test@test.nl",
		Password: "ABC123PXPXdddd@",
	})

	messages := []*message.Message{
		{
			ID:              1,
			UID:             "<to be replaced>",
			Text:            "lorum ipsum",
			CreatedAt:       time.Now(),
			CreatedByUserID: user.ID,
			User: &message.User{
				ID:       user.ID,
				UID:      user.UID,
				Username: user.Username,
			},
		},
	}

	for i, msg := range messages {
		m, err := store.Message.Create(message.CreateMessage{
			UID:             msg.UID,
			Text:            msg.Text,
			CreatedByUserID: msg.CreatedByUserID,
			CreatedAt:       msg.CreatedAt,
		})

		if err != nil {
			assert.Fail(t, err.Error())
		}

		messages[i].UID = m.UID
		messages[i].CreatedAt = m.CreatedAt
	}

	req, err := http.NewRequest("DELETE", fmt.Sprintf("/%s", messages[0].UID), nil)
	assert.NoError(t, err)

	token, _ := auth.CreateAccessToken(user.UID, "1", time.Now().Add(1000000).Unix())
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	rr := httptest.NewRecorder()
	NewMessageRouter(store).ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))

	expectedResponds := mapMessage(messages[0])

	// Check the response body is what we expect.
	expected, _ := json.Marshal(expectedResponds)
	assert.Equal(t, string(expected), rr.Body.String(), "handler returned unexpected body")
}

func TestUnauthorizedDeleteMessage(t *testing.T) {
	store := domain.NewInMemoryStore(memstore.NewStore())
	user1, _ := store.User.Create(user.CreateUser{
		Username: "test",
		Email:    "test@test.nl",
		Password: "ABC123PXPXdddd@",
	})

	messages := []*message.Message{
		{
			ID:              1,
			UID:             "<to be replaced>",
			Text:            "lorum ipsum",
			CreatedAt:       time.Now(),
			CreatedByUserID: user1.ID,
			User: &message.User{
				ID:       user1.ID,
				UID:      user1.UID,
				Username: user1.Username,
			},
		},
	}

	user2, _ := store.User.Create(user.CreateUser{
		Username: "test2",
		Email:    "test2@test.nl",
		Password: "ABC123PXPXdddd@",
	})

	for i, msg := range messages {
		m, err := store.Message.Create(message.CreateMessage{
			UID:             msg.UID,
			Text:            msg.Text,
			CreatedByUserID: msg.CreatedByUserID,
			CreatedAt:       msg.CreatedAt,
		})

		if err != nil {
			assert.Fail(t, err.Error())
		}

		messages[i].UID = m.UID
		messages[i].CreatedAt = m.CreatedAt
	}

	req, err := http.NewRequest("DELETE", fmt.Sprintf("/%s", messages[0].UID), nil)
	assert.NoError(t, err)

	token, _ := auth.CreateAccessToken(user2.UID, "1", time.Now().Add(1000000).Unix())
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	rr := httptest.NewRecorder()
	NewMessageRouter(store).ServeHTTP(rr, req)

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
	store := domain.NewInMemoryStore(memstore.NewStore())
	u, _ := store.User.Create(user.CreateUser{
		Username: "test",
		Email:    "test@test.nl",
		Password: "ABC123PXPXdddd@",
	})

	req, err := http.NewRequest("DELETE", "/abcdefg1234", nil)
	assert.NoError(t, err)

	token, _ := auth.CreateAccessToken(u.UID, "1", time.Now().Add(1000000).Unix())
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	rr := httptest.NewRecorder()
	NewMessageRouter(store).ServeHTTP(rr, req)

	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))
	assert.Equal(t, http.StatusNotFound, rr.Code)

	expectedResponds := response.ErrorResponds{
		Error:   http.StatusText(http.StatusNotFound),
		Message: "resource not found or unavailable",
	}
	// Check the response body is what we expect.
	expected, _ := json.Marshal(expectedResponds)
	assert.Equal(t, string(expected), rr.Body.String(), rr.Body.String())
}
