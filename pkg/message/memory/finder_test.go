package memory

import (
	"fmt"
	"testing"
	"time"

	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/dwethmar/atami/pkg/message"
	"github.com/dwethmar/atami/pkg/message/memory/util"
)

func generateTestMessages(size int) []message.CreateRequest {
	messages := make([]message.CreateRequest, size)
	for i := 0; i < size; i++ {
		messages[i] = message.CreateRequest{
			Text:            fmt.Sprintf("Lorum ipsum %v", i+1),
			CreatedByUserID: 1,
		}
	}
	return messages
}

func setup() (*memstore.Store, []message.Message) {
	store := memstore.NewStore()
	util.AddTestUser(store, 1)

	service := New(store)
	msgs := make([]message.Message, 100)
	for i, newMsg := range generateTestMessages(100) {
		if msg, err := service.Create(newMsg); err == nil {
			msgs[i] = *msg
		} else {
			fmt.Printf("error: %s", err)
			panic(1)
		}
	}
	return store, msgs
}

func TestByUID(t *testing.T) {
	store, messages := setup()
	message.TestFindByUID(t, NewFinder(store), messages[0].UID, messages[0])
}

func TestFindByID(t *testing.T) {
	store, _ := setup()
	message.TestFindByID(t, NewFinder(store), 1, message.Message{
		ID:              1,
		UID:             "abcdef",
		Text:            "Lorum ipsum 1",
		CreatedByUserID: 1,
		CreatedAt:       time.Now(),
	})
}

func TestNotFound(t *testing.T) {
	store, _ := setup()
	message.TestNotFound(t, NewFinder(store))
}

func TestFindAll(t *testing.T) {
	store, messages := setup()
	message.TestFind(t, NewFinder(store), 100, messages)
}
