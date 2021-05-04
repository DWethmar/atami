package memory

import (
	"fmt"
	"testing"
	"time"

	"github.com/dwethmar/atami/pkg/domain/message"
	"github.com/dwethmar/atami/pkg/domain/message/memory/util"
	"github.com/dwethmar/atami/pkg/domain/message/test"
	"github.com/dwethmar/atami/pkg/memstore"
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
	store := memstore.NewStore()
	util.AddTestUser(store, 1)

	creator := NewCreator(store)
	msgs := make([]message.Message, 100)
	for i, newMsg := range generateTestMessages(100) {
		if msg, err := creator.Create(newMsg); err == nil {
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
	test.FindByUID(t, NewFinder(store), messages[0].UID, messages[0])
}

func TestFindByID(t *testing.T) {
	store, _ := setup()
	test.FindByID(t, NewFinder(store), 1, message.Message{
		ID:              1,
		UID:             "abcdef",
		Text:            "Lorum ipsum 1",
		CreatedByUserID: 1,
		CreatedAt:       time.Now(),
	})
}

func TestNotFound(t *testing.T) {
	store, _ := setup()
	test.NotFound(t, NewFinder(store))
}

func TestFindAll(t *testing.T) {
	store, messages := setup()
	test.Find(t, NewFinder(store), 100, messages)
}
