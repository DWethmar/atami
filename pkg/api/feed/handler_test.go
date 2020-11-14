package feed

import (
	"testing"

	"github.com/dwethmar/atami/pkg/auth"
	"github.com/dwethmar/atami/pkg/service"
)

func TestNewHandler(t *testing.T) {
	authService := service.NewAuthServiceMemory()
	authService.Register(auth.CreateUser{
		Username: "lol",
	})
}
