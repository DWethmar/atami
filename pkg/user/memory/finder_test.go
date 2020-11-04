package memory

import (
	"fmt"
	"testing"
	"time"

	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/dwethmar/atami/pkg/user"
	"github.com/stretchr/testify/assert"
)

func TestFindAll(t *testing.T) {
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

	finder := NewFinder(store)
	user.TestFindAll(t, finder, 2, []user.User{
		a,
		b,
	})
}

func TestFindByID(t *testing.T) {
	store := memstore.New()
	a := user.User{
		ID:        1,
		UID:       "x",
		Email:     "test@test.nl",
		CreatedAt: time.Now(),
	}
	assert.True(t, store.Add(a.ID.String(), a))

	finder := NewFinder(store)
	user.TestFindByID(t, finder, 1)
}

func TestNotFound(t *testing.T) {
	store := memstore.New()
	finder := NewFinder(store)
	user.TestNotFound(t, finder)
}

func TestFindByEmail(t *testing.T) {
	store := memstore.New()

	for i := 1; i < 100; i++ {
		a := user.User{
			ID:        user.ID(i),
			UID:       user.UID(fmt.Sprintf("%b", i)),
			Email:     fmt.Sprintf("test-%d@test.com", i),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		assert.True(t, store.Add(a.ID.String(), a))
	}

	finder := NewFinder(store)
	user.TestFindByEmail(t, finder, "test-44@test.com")
}
