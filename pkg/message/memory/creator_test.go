package memory

import (
	"testing"

	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/dwethmar/atami/pkg/message"
)

func TestCreate(t *testing.T) {
	newMessage := message.CreateMessageRequest{
		Text:            "lorum ipsum",
		CreatedByUserID: 1,
	}
	message.TestCreator(t, NewCreator(memstore.New()), newMessage)
}
