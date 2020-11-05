package memory

import (
	"fmt"
	"testing"
	"time"

	"github.com/dwethmar/atami/pkg/auth"
	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/stretchr/testify/assert"
)

func TestFindAll(t *testing.T) {
	store := memstore.New()
	a := auth.User{
		ID:        1,
		UID:       "x",
		Email:     "test@test.nl",
		CreatedAt: time.Now(),
	}
	assert.True(t, store.Add(a.ID.String(), a))

	b := auth.User{
		ID:        2,
		UID:       "y",
		Email:     "test2@test.nl",
		CreatedAt: time.Now(),
	}
	assert.True(t, store.Add(b.ID.String(), b))

	finder := NewFinder(store)
	auth.TestFindAll(t, finder, 2, []auth.User{
		a,
		b,
	})
}

func TestFindByID(t *testing.T) {
	store := memstore.New()
	a := auth.User{
		ID:        1,
		UID:       "x",
		Email:     "test@test.nl",
		CreatedAt: time.Now(),
	}
	assert.True(t, store.Add(a.ID.String(), a))

	finder := NewFinder(store)
	auth.TestFindByID(t, finder, 1)
}

func TestNotFound(t *testing.T) {
	store := memstore.New()
	finder := NewFinder(store)
	auth.TestNotFound(t, finder)
}

func TestFindByEmail(t *testing.T) {
	store := memstore.New()

	for i := 1; i < 100; i++ {
		a := auth.User{
			ID:        auth.ID(i),
			UID:       auth.UID(fmt.Sprintf("%b", i)),
			Username:  fmt.Sprintf("username-%d", i),
			Email:     fmt.Sprintf("test-%d@test.com", i),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		assert.True(t, store.Add(a.ID.String(), a))
	}

	finder := NewFinder(store)
	auth.TestFindByEmail(t, finder, "test-44@test.com")
}

func TestFindByUsername(t *testing.T) {
	store := memstore.New()

	for i := 1; i < 100; i++ {
		a := auth.User{
			ID:        auth.ID(i),
			UID:       auth.UID(fmt.Sprintf("%b", i)),
			Username:  fmt.Sprintf("username-%d", i),
			Email:     fmt.Sprintf("test-%d@test.com", i),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		// fmt.Printf("Username: %v\n", a.Username)

		assert.True(t, store.Add(a.ID.String(), a))
	}

	finder := NewFinder(store)
	auth.TestFindByUsername(t, finder, "username-44")
}
