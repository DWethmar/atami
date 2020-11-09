package memory

import (
	"testing"

	"github.com/dwethmar/atami/pkg/auth"
	"github.com/dwethmar/atami/pkg/memstore"
)

func TestCreate(t *testing.T) {
	s := memstore.New()
	register := NewRegistrator(NewFinder(s), auth.NewDefaultValidator(), s)
	auth.TestRegister(t, register, auth.CreateUser{
		Username: "username",
		Email:    "test@test.nl",
		Password: "!Test123",
	})
}

func TestDuplicateUsername(t *testing.T) {
	newUser := auth.CreateUser{
		Username: "username",
		Email:    "test@test.nl",
		Password: "!Test123",
	}
	s := memstore.New()
	register := NewRegistrator(NewFinder(s), auth.NewDefaultValidator(), s)
	auth.TestDuplicateUsername(t, register, newUser)
}

func TestDuplicateEmail(t *testing.T) {
	newUser := auth.CreateUser{
		Username: "username",
		Email:    "test@test.nl",
		Password: "!Test123",
	}
	s := memstore.New()
	register := NewRegistrator(NewFinder(s), auth.NewDefaultValidator(), s)
	auth.TestDuplicateEmail(t, register, newUser)
}

// TestEmptyPassword test if the correct error is returned
func TestEmptyPassword(t *testing.T) {
	s := memstore.New()
	register := NewRegistrator(NewFinder(s), auth.NewDefaultValidator(), s)
	auth.TestEmptyPassword(t, register)
}
