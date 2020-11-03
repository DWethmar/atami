package memory

import (
	"testing"
	"time"

	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/dwethmar/atami/pkg/message"
	"github.com/stretchr/testify/assert"
)

func TestReadOne(t *testing.T) {
	store := memstore.New()
	a := message.Message{
		ID:        1,
		UID:       "x",
		Content:   "sd1",
		CreatedAt: time.Now(),
	}
	assert.True(t, store.Add(a.ID.String(), a))

	reader := NewReader(store)
	message.TestReadOne(t, reader, a.ID, a)
}

func TestNotFound(t *testing.T) {
	store := memstore.New()
	reader := NewReader(store)
	message.TestNotFound(t, reader)
}

func TestReadAll(t *testing.T) {
	store := memstore.New()
	a := message.Message{
		ID:        1,
		UID:       "x",
		Content:   "sd1",
		CreatedAt: time.Now(),
	}
	assert.True(t, store.Add(a.ID.String(), a))

	b := message.Message{
		ID:        2,
		UID:       "y",
		Content:   "sd2",
		CreatedAt: time.Now(),
	}
	assert.True(t, store.Add(b.ID.String(), b))

	reader := NewReader(store)
	message.TestReadAll(t, reader, 2, []message.Message{
		a,
		b,
	})
}
