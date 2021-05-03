package memory

import (
	"testing"

	"github.com/dwethmar/atami/pkg/domain/message"
	"github.com/dwethmar/atami/pkg/domain/message/memory/util"
	"github.com/dwethmar/atami/pkg/domain/message/test"
	"github.com/dwethmar/atami/pkg/memstore"
)

func TestCreate(t *testing.T) {
	store := memstore.NewStore()
	util.AddTestUser(store, 1)
	newMessage := message.CreateMessage{
		Text:            "lorum ipsum",
		CreatedByUserID: 1,
	}
	test.Create(t, NewCreator(store), newMessage)
}

func TestInvalidCreate(t *testing.T) {
	store := memstore.NewStore()
	util.AddTestUser(store, 1)
	newMessage := message.CreateMessage{
		Text:            "lorum ipsum",
	}
	test.InvalidCreate(t, NewCreator(store), newMessage)
}
