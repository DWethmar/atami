package memory

import (
	"testing"

	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/dwethmar/atami/pkg/user"
	"github.com/dwethmar/atami/pkg/user/test"
)

func TestCreate(t *testing.T) {
	s := memstore.NewStore()
	register := NewCreator(s)
	test.TestCreator(t, register, user.CreateUser{
		Username: "username",
		Email:    "test@test.nl",
		Password: "!Test123",
	})
}

func TestDuplicateUsername(t *testing.T) {
	newUser := user.CreateUser{
		Username: "username",
		Email:    "test@test.nl",
		Password: "!Test123",
	}
	s := memstore.NewStore()
	register := NewCreator(s)
	test.TestDuplicateUsername(t, register, newUser)
}

func TestDuplicateEmail(t *testing.T) {
	newUser := user.CreateUser{
		Username: "username",
		Email:    "test@test.nl",
		Password: "!Test123",
	}
	s := memstore.NewStore()
	register := NewCreator(s)
	test.TestDuplicateEmail(t, register, newUser)
}

// TestEmptyPassword test if the correct error is returned
func TestEmptyPassword(t *testing.T) {
	s := memstore.NewStore()
	register := NewCreator(s)
	test.TestEmptyPassword(t, register)
}
