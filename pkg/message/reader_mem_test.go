package message

import (
	"testing"
	"time"

	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/stretchr/testify/assert"
)

func TestReadOne(t *testing.T) {
	store := memstore.NewMemstore()
	a := Message{
		ID:        1,
		Content:   "sd1",
		CreatedAt: time.Now(),
	}
	assert.True(t, store.Add("1", a))

	reader := NewMemReader(store)
	testReadOne(t, *reader, 1, a)
}

func TestReadAll(t *testing.T) {
	store := memstore.NewMemstore()
	a := Message{
		ID:        1,
		Content:   "sd1",
		CreatedAt: time.Now(),
	}
	assert.True(t, store.Add("1", a))

	b := Message{
		ID:        2,
		Content:   "sd2",
		CreatedAt: time.Now(),
	}
	assert.True(t, store.Add("2", b))

	reader := NewMemReader(store)
	testReadAll(t, *reader, 2, []Message{
		a, b,
	})
}
