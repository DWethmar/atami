package memory

import (
	"fmt"
	"testing"

	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/dwethmar/atami/pkg/message"
)

func generateTestMessages(size int) []message.CreateMessage {
	messages := make([]message.CreateMessage, size)
	for i := 0; i < size; i++ {
		messages[i] = message.CreateMessage{
			Text:            fmt.Sprintf("Lorum ipsum %v", i+1),
			CreatedByUserID: 1,
		}
	}
	return messages
}

func setup() (*memstore.Store, []message.Message) {
	store := memstore.New()
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

func TestReadOne(t *testing.T) {
	store, _ := setup()
	message.TestFindOne(t, NewFinder(store), 1, message.Message{
		ID:              1,
		UID:             "",
		Text:            "Lorum ipsum 1",
		CreatedByUserID: 1,
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
