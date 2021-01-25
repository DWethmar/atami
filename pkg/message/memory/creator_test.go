package memory

import (
	"testing"

	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/dwethmar/atami/pkg/message"
	"github.com/dwethmar/atami/pkg/message/memory/util"
)

func TestCreate(t *testing.T) {
	store := memstore.NewStore()
	util.AddTestUser(store, 1)
	newMessage := message.CreateRequest{
		Text:            "lorum ipsum",
		CreatedByUserID: 1,
	}
	message.TestCreator(t, NewCreator(store), newMessage)
}
