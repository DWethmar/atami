package memory

import (
	"testing"

	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/dwethmar/atami/pkg/user"
)

func TestCreate(t *testing.T) {
	s := memstore.NewStore()
	register := NewCreator(s)
	user.TestCreator(t, register, user.CreateUserRequest{
		Username: "username",
		Email:    "test@test.nl",
		Password: "!Test123",
	})
}

func TestDuplicateUsername(t *testing.T) {
	newUser := user.CreateUserRequest{
		Username: "username",
		Email:    "test@test.nl",
		Password: "!Test123",
	}
	s := memstore.NewStore()
	register := NewCreator(s)
	user.TestDuplicateUsername(t, register, newUser)
}

func TestDuplicateEmail(t *testing.T) {
	newUser := user.CreateUserRequest{
		Username: "username",
		Email:    "test@test.nl",
		Password: "!Test123",
	}
	s := memstore.NewStore()
	register := NewCreator(s)
	user.TestDuplicateEmail(t, register, newUser)
}

// TestEmptyPassword test if the correct error is returned
func TestEmptyPassword(t *testing.T) {
	s := memstore.NewStore()
	register := NewCreator(s)
	user.TestEmptyPassword(t, register)
}
