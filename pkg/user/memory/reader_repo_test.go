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

	repo := NewReaderRepository(store)
	user.TestReadOne(t, repo, 1, a)
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

	repo := NewReaderRepository(store)
	user.TestReadAll(t, user.NewReader(repo), 2, []user.User{
		a,
		b,
	})
}
