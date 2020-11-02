package memory

import (
	"testing"

	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/dwethmar/atami/pkg/message"
)

func TestCreate(t *testing.T) {
	newMessage := message.NewMessage{
		Content: "wow",
	}

	repo := NewCreatorRepository(memstore.New())
	message.TestCreator(t, repo, newMessage)
}
