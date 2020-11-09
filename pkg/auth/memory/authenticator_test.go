package memory

import (
	"testing"

	"github.com/dwethmar/atami/pkg/auth"
	"github.com/dwethmar/atami/pkg/memstore"
)

func TestAuthenticate(t *testing.T) {
	store := memstore.New()
	registrator := NewRegistrator(NewFinder(store), auth.NewDefaultValidator(), store)

	registrator.Register(auth.CreateUser{
		Username: "test",
		Email:    "test@test.com",
		Password: "TestPassw0rd0987!@#s",
	})

	authenticator := NewAuthenticator(store)
	auth.TestAuthenticate(t, authenticator, auth.Credentials{
		Email:    "test@test.com",
		Password: "TestPassw0rd0987!@#s",
	})
}
