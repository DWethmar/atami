package memory

import (
	"testing"
	"time"

	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/dwethmar/atami/pkg/message"
	"github.com/stretchr/testify/assert"
)

func TestDelete(t *testing.T) {
	store := memstore.New()
	a := message.Message{
		ID:        1,
		UID:       "x",
		Content:   "sd1",
		CreatedAt: time.Now(),
	}
	assert.True(t, store.Add(a.ID.String(), a))

	repo := NewReaderRepository(store)
	message.TestReadOne(t, repo, a.ID, a)
}
