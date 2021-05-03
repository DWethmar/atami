package memory

import (
	"testing"
	"time"

	"github.com/dwethmar/atami/pkg/domain/message"
	"github.com/dwethmar/atami/pkg/domain/message/memory/util"
	"github.com/dwethmar/atami/pkg/domain/message/test"
	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/stretchr/testify/assert"
)

func TestDelete(t *testing.T) {
	store := memstore.NewStore()
	a := message.Message{
		ID:        1,
		UID:       "x",
		Text:      "sd1",
		CreatedAt: time.Now(),
	}
	assert.True(t, store.GetMessages().Put(a.ID, util.ToMemory(a)))

	deleter := NewDeleter(store)
	test.Delete(t, deleter, a.ID)
}
