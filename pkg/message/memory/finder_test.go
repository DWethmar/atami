package memory

import (
	"fmt"
	"testing"

	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/dwethmar/atami/pkg/message"
)

func generateTestMessages(size int) []message.NewMessage {
	messages := make([]message.NewMessage, size)
	for i := 0; i < size; i++ {
		messages[i] = message.NewMessage{
			Text:            fmt.Sprintf("Lorum ipsum %v", i+1),
			CreatedByUserID: 1,
		}
	}
	return messages
}

func setup() (*message.Finder, []message.Message) {
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
	return NewFinder(store), msgs
}

func TestReadOne(t *testing.T) {
	reader, _ := setup()
	message.TestFindOne(t, reader, 1, message.Message{
		ID:              1,
		UID:             "",
		Text:            "Lorum ipsum 1",
		CreatedByUserID: 1,
	})
}

func TestNotFound(t *testing.T) {
	finder, _ := setup()
	message.TestNotFound(t, finder)
}

func TestReadAll(t *testing.T) {
	finder, messages := setup()
	message.TestFindAll(t, finder, 100, messages)
}
