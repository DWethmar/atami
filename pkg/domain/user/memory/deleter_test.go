package memory

import (
	"testing"
	"time"

	"github.com/dwethmar/atami/pkg/domain/user"
	"github.com/dwethmar/atami/pkg/domain/user/test"
	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/stretchr/testify/assert"
)

func TestDelete(t *testing.T) {
	memstore := memstore.NewStore()
	creator := NewCreator(memstore)
	user, err := creator.Create(user.CreateUser{
		UID:       "UID1",
		Username:  "test",
		Email:     "test1d@test.nl",
		Password:  "kjashdkljhasd@@88ssKK",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err != nil {
		assert.Fail(t, err.Error())
	}

	deleter := NewDeleter(memstore)
	test.TestDelete(t, deleter, user.ID)
}
