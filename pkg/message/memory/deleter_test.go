package memory

import (
	"strconv"
	"testing"
	"time"

	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/dwethmar/atami/pkg/message"
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
	assert.True(t, store.GetMessages().Put(strconv.Itoa(a.ID), a))

	deleter := NewDeleter(store)
	message.TestDelete(t, deleter, a.ID)
}
