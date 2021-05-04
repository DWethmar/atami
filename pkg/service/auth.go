package service

import (
	"github.com/dwethmar/atami/pkg/auth"
	"github.com/dwethmar/atami/pkg/domain/user"
)

// NewAuthServicePostgres create a new postgres auth service
func NewAuthService(finder *user.Finder, creator *user.Creator) *auth.Service {
	return auth.NewService(
		*auth.NewAuthenticator(finder),
		*auth.NewRegistrator(creator, finder),
	)
}
