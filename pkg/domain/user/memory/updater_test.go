package memory

import (
	"testing"

	"github.com/dwethmar/atami/pkg/domain/user"
	"github.com/dwethmar/atami/pkg/domain/user/test"
	"github.com/dwethmar/atami/pkg/memstore"
)

func TestUpdate(t *testing.T) {
	memstore := memstore.NewStore()
	creator := NewCreator(memstore, NewFinder(memstore))

	creator.Create(user.CreateUser{
		Username: "test",
		Email:    "test@test.nl",
		Password: "Passwordz@1",
	})

	test.TestUpdater(t, NewUpdater(memstore), user.UpdateUser{
		Biography: "lorum ipsum",
	})
}
