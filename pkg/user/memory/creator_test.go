package memory

import (
	"testing"

	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/dwethmar/atami/pkg/user"
	"github.com/dwethmar/atami/pkg/validate"
)

var validator = user.NewValidator(validate.NewEmailValidator())

func TestCreate(t *testing.T) {
	creator := NewCreator(validator, memstore.New())
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
	creator := NewCreator(validator, memstore.New())
	user.TestDuplicateEmail(t, creator, newUser)
}

// TestEmptyPassword test if the correct error is returned
func TestEmptyPassword(t *testing.T) {
	creator := NewCreator(validator, memstore.New())
	user.TestEmptyPassword(t, creator)
}
