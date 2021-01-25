package memory

import (
	"testing"

	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/dwethmar/atami/pkg/user"
)

func TestUpdate(t *testing.T) {
	store := memstore.NewStore()
	updater := NewUpdater(store)

	user.TestUpdater(t, updater, user.UpdateRequest{
		Biography: "lorum ipsum",
	})
}
