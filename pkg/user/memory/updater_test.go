package memory

import (
	"testing"

	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/dwethmar/atami/pkg/user"
)

func TestUpdate(t *testing.T) {
	store := memstore.NewStore()
	updater := NewUpdater(store)

	register := NewCreator(store)
	register.Create(user.CreateRequest{
		Username: "test",
		Email:    "test@test.nl",
		Password: "Passwordz@1",
	})

	user.TestUpdater(t, updater, user.UpdateRequest{
		Biography: "lorum ipsum",
	})
}
