package memory

import (
	"strconv"
	"testing"
	"time"

	"github.com/dwethmar/atami/pkg/auth"
	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/stretchr/testify/assert"
)

func TestDelete(t *testing.T) {
	store := memstore.New()
	a := userRecord{
		ID:        1,
		UID:       "x",
		Email:     "test@test.nl",
		CreatedAt: time.Now(),
	}
	assert.True(t, store.Add(strconv.Itoa(a.ID), a))

	deleter := NewDeleter(store)
	auth.TestDelete(t, deleter, a.ID)
}
