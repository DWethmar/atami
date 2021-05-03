package memory

import (
	"testing"

	"github.com/dwethmar/atami/pkg/domain/user"
	"github.com/dwethmar/atami/pkg/domain/user/test"
	"github.com/dwethmar/atami/pkg/memstore"
)

func TestUpdate(t *testing.T) {
	store := memstore.NewStore()
	updater := NewUpdater(store)

	register := NewCreator(store)
	register.Create(user.CreateUser{
		Username: "test",
		Email:    "test@test.nl",
		Password: "Passwordz@1",
	})

	test.TestUpdater(t, updater, user.UpdateUser{
		Biography: "lorum ipsum",
	})
}