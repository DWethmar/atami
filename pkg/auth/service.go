package auth

import "github.com/dwethmar/atami/pkg/domain/user"

// Service defines interations with users
type Service struct {
	*Authenticator
}

// CreateService creates a new auth service
func CreateService(
	a *Authenticator,
) *Service {
	return &Service{
		Authenticator: a,
	}
}

// NewService create a new postgres auth service
func NewService(finder *user.Finder) *Service {
	return CreateService(
		NewAuthenticator(finder),
	)
}
