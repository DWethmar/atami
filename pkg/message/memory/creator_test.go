package memory

import (
	"testing"

	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/dwethmar/atami/pkg/message"
)

func TestCreate(t *testing.T) {
	newMessage := message.CreateMessage{
		Text: "lorum ipsum",
	}
	message.TestCreator(t, NewCreator(memstore.New()), newMessage)
}
