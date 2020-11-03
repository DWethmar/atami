package memory

import (
	"testing"
	"time"

	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/dwethmar/atami/pkg/user"
	"github.com/stretchr/testify/assert"
)

func TestReadOne(t *testing.T) {
	store := memstore.New()
	a := user.User{
		ID:        1,
		UID:       "x",
		Email:     "test@test.nl",
		CreatedAt: time.Now(),
	}
	assert.True(t, store.Add(a.ID.String(), a))

	reader := NewReader(store)
	user.TestReadOne(t, reader, 1, a)
}

func TestNotFound(t *testing.T) {
	store := memstore.New()
	reader := NewReader(store)
	user.TestNotFound(t, reader)
}

func TestReadAll(t *testing.T) {
	store := memstore.New()
	a := user.User{
		ID:        1,
		UID:       "x",
		Email:     "test@test.nl",
		CreatedAt: time.Now(),
	}
	assert.True(t, store.Add(a.ID.String(), a))

	b := user.User{
		ID:        2,
		UID:       "y",
		Email:     "test2@test.nl",
		CreatedAt: time.Now(),
	}
	assert.True(t, store.Add(b.ID.String(), b))

	reader := NewReader(store)
	user.TestReadAll(t, reader, 2, []user.User{
		a,
		b,
	})
}
