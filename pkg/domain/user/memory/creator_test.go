package memory

import (
	"testing"

	"github.com/dwethmar/atami/pkg/domain/user"
	"github.com/dwethmar/atami/pkg/domain/user/test"
	"github.com/dwethmar/atami/pkg/memstore"
)

func TestCreate(t *testing.T) {
	s := memstore.NewStore()
	creator := NewCreator(s, NewFinder(s))

	test.Creator(t, creator, user.CreateUser{
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
	creator := NewCreator(s, NewFinder(s))
	test.DuplicateUsername(t, creator, newUser)
}

func TestDuplicateEmail(t *testing.T) {
	newUser := user.CreateUser{
		Username: "username",
		Email:    "test@test.nl",
		Password: "!Test123",
	}
	s := memstore.NewStore()
	creator := NewCreator(s, NewFinder(s))
	test.DuplicateEmail(t, creator, newUser)
}

// TestEmptyPassword test if the correct error is returned
func TestEmptyPassword(t *testing.T) {
	s := memstore.NewStore()
	creator := NewCreator(s, NewFinder(s))
	test.EmptyPassword(t, creator)
}
