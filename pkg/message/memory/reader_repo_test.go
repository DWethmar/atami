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
		Content:   "sd1",
		CreatedAt: time.Now(),
	}
	assert.True(t, store.Add("1", a))

	repo := NewReaderRepository(store)
	message.TestReadOne(t, repo, 1, a)
}

func TestReadAll(t *testing.T) {
	store := memstore.New()
	a := message.Message{
		ID:        1,
		Content:   "sd1",
		CreatedAt: time.Now(),
	}
	assert.True(t, store.Add("1", a))

	b := message.Message{
		ID:        2,
		Content:   "sd2",
		CreatedAt: time.Now(),
	}
	assert.True(t, store.Add("2", b))

	repo := NewReaderRepository(store)
	message.TestReadAll(t, message.NewReader(repo), 2, []message.Message{
		a, b,
	})
}
