package memory

import (
	"testing"

	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/dwethmar/atami/pkg/node"
	"github.com/dwethmar/atami/pkg/node/memory/util"
)

func TestCreate(t *testing.T) {
	store := memstore.NewStore()
	util.AddTestUser(store, 1)
	newNode := node.CreateRequest{
		Text:            "lorum ipsum",
		CreatedByUserID: 1,
	}
	node.TestCreator(t, NewCreator(store), newNode)
}
