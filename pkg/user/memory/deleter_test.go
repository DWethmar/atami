package memory

import (
	"testing"

	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/dwethmar/atami/pkg/user"
	"github.com/dwethmar/atami/pkg/user/memory/util"
)

func TestDelete(t *testing.T) {
	store := memstore.NewStore()
	util.AddTestUser(store, 1)
	deleter := NewDeleter(store)
	user.TestDelete(t, deleter, 1)
}
