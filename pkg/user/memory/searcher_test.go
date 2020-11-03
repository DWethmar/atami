package memory

import (
	"fmt"
	"testing"
	"time"

	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/dwethmar/atami/pkg/user"
	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	store := memstore.New()

	for i := 1; i < 100; i++ {
		a := user.User{
			ID:        user.ID(i),
			UID:       user.UID(fmt.Sprintf("UID%b", i)),
			Email:     fmt.Sprintf("test-%b@test.nl", i),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		if i == 50 {
			a.Email = "hit@testr.nl"
		}

		assert.True(t, store.Add(a.ID.String(), a))
	}

	searcher := NewSearcher(store)
	user.TestSearchByEmail(t, searcher, 1, "hit@testr.nl")
	user.TestSearchByEmail(t, searcher, 0, "miss@testr.nl")
}
