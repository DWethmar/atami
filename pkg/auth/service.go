package auth

import "github.com/dwethmar/atami/pkg/domain/user"

// Service defines interations with users
type Service struct {
	*Authenticator
	*Registrator
	*Validator
}

// CreateService creates a new auth service
func CreateService(
	a *Authenticator,
	r *Registrator,
	v *Validator,
) *Service {
	return &Service{
		Authenticator: a,
		Registrator:   r,
		Validator:     v,
	}
}

// NewService create a new postgres auth service
func NewService(finder *user.Finder, creator *user.Creator) *Service {
	return CreateService(
		NewAuthenticator(finder),
		NewRegistrator(creator, finder),
		NewDefaultValidator(),
	)
}
