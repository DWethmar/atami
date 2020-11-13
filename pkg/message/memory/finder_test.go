package memory

import (
	"fmt"
	"testing"

	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/dwethmar/atami/pkg/message"
	"github.com/dwethmar/atami/pkg/model"
)

func generateTestMessages(size int) []message.NewMessage {
	messages := make([]message.NewMessage, size)
	for i := 0; i < size; i++ {
		messages[i] = message.NewMessage{
			Text:      fmt.Sprintf("Lorum ipsum %v", i+1),
			CreatedBy: model.UserID(1),
		}
	}
	return messages
}

func setup() (*message.Finder, []model.Message) {
	store := memstore.New()
	service := NewService(store)
	msgs := make([]model.Message, 100)
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
	message.TestFindOne(t, reader, model.MessageID(1), model.Message{
		ID:        1,
		UID:       "",
		Text:      "Lorum ipsum 1",
		CreatedBy: model.UserID(1),
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
