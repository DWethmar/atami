package memory

import (
	"testing"

	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/dwethmar/atami/pkg/message"
	"github.com/dwethmar/atami/pkg/message/memory/util"
	"github.com/dwethmar/atami/pkg/message/test"
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
