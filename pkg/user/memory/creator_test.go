package memory

import (
	"testing"

	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/dwethmar/atami/pkg/user"
)

func TestCreate(t *testing.T) {
	creator := NewCreator(memstore.New())
	user.TestCreator(t, creator, user.NewUser{
		Username: "username",
		Email:    "test@test.nl",
		Password: "test123",
	})
}

func TestDuplicateEmail(t *testing.T) {
	newUser := user.NewUser{
		Username: "username",
		Email:    "test@test.nl",
		Password: "test123",
	}
	creator := NewCreator(memstore.New())
	user.TestDuplicateEmail(t, creator, newUser)
}

// TestEmptyPassword test if the correct error is returned
func TestEmptyPassword(t *testing.T) {
	creator := NewCreator(memstore.New())
	user.TestEmptyPassword(t, creator)
}
