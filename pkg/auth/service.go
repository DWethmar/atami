package auth

// Service defines interations with users
type Service struct {
	Authenticator
	Registrator
	Validator
}

// NewService creates a new user service
func NewService(
	a Authenticator,
	r Registrator,
) *Service {
	return &Service{
		Authenticator: a,
		Registrator:   r,
		Validator:     *NewDefaultValidator(),
	}
}
