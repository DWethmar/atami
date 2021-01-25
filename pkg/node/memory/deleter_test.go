package memory

import (
	"testing"
	"time"

	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/dwethmar/atami/pkg/node"
	"github.com/dwethmar/atami/pkg/node/memory/util"
	"github.com/stretchr/testify/assert"
)

func TestDelete(t *testing.T) {
	store := memstore.NewStore()
	a := node.Node{
		ID:        1,
		UID:       "x",
		Text:      "sd1",
		CreatedAt: time.Now(),
	}
	assert.True(t, store.GetNodes().Put(a.ID, util.ToMemory(a)))

	deleter := NewDeleter(store)
	node.TestDelete(t, deleter, a.ID)
}
