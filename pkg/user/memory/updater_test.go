package memory

import (
	"testing"

	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/dwethmar/atami/pkg/user"
	"github.com/dwethmar/atami/pkg/user/test"
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

	test.TestUpdater(t, updater, user.UpdateRequest{
		Biography: "lorum ipsum",
	})
}
