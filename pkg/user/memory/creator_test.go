package memory

import (
	"testing"

	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/dwethmar/atami/pkg/user"
)

func TestCreate(t *testing.T) {
	s := memstore.New()
	register := NewCreator(user.NewDefaultValidator(), s)
	user.TestCreator(t, register, user.CreateUser{
		Username:       "username",
		Email:          "test@test.nl",
		HashedPassword: "!Test123",
	})
}

func TestDuplicateUsername(t *testing.T) {
	newUser := user.CreateUser{
		Username:       "username",
		Email:          "test@test.nl",
		HashedPassword: "!Test123",
	}
	s := memstore.New()
	register := NewCreator(user.NewDefaultValidator(), s)
	user.TestDuplicateUsername(t, register, newUser)
}

func TestDuplicateEmail(t *testing.T) {
	newUser := user.CreateUser{
		Username:       "username",
		Email:          "test@test.nl",
		HashedPassword: "!Test123",
	}
	s := memstore.New()
	register := NewCreator(user.NewDefaultValidator(), s)
	user.TestDuplicateEmail(t, register, newUser)
}

// TestEmptyPassword test if the correct error is returned
func TestEmptyPassword(t *testing.T) {
	s := memstore.New()
	register := NewCreator(user.NewDefaultValidator(), s)
	user.TestEmptyPassword(t, register)
}
