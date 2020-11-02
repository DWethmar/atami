package memory

import (
	"testing"

	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/dwethmar/atami/pkg/user"
)

func TestCreate(t *testing.T) {
	repo := NewCreatorRepository(memstore.New())
	user.TestCreator(t, repo, user.NewUser{
		Email:    "test@test.nl",
		Password: "test123",
	})
}

func TestDuplicateEmail(t *testing.T) {
	newUser := user.NewUser{
		Username: "asd",
		Email:    "test@test.nl",
		Password: "test123",
	}
	repo := NewCreatorRepository(memstore.New())
	user.TestDuplicateEmail(t, repo, newUser)
}

// TestEmptyPassword test if the correct error is returned
func TestEmptyPassword(t *testing.T) {
	repo := NewCreatorRepository(memstore.New())
	user.TestEmptyPassword(t, repo)
}
