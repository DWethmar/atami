package auth

// Service defines interations with users
type Service struct {
	Authenticator
	Finder
	Deleter
	Registrator
	Validator
}

// NewService creates a new user service
func NewService(
	a Authenticator,
	f Finder,
	d Deleter,
	r Registrator,
	v Validator,
) *Service {
	return &Service{
		Authenticator: a,
		Finder:        f,
		Deleter:       d,
		Registrator:   r,
		Validator:     v,
	}
}
