package memory

import (
	"testing"

	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/dwethmar/atami/pkg/user/memory/util"
	"github.com/dwethmar/atami/pkg/user/test"
)

func TestDelete(t *testing.T) {
	store := memstore.NewStore()
	util.AddTestUser(store, 1)
	deleter := NewDeleter(store)
	test.TestDelete(t, deleter, 1)
}
