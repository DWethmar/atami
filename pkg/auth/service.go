package auth

// Service defines interations with users
type Service interface {
	Authenticate(credentials Credentials) (bool, error)
	FindAll() ([]*User, error)
	FindByID(ID ID) (*User, error)
	FindByEmail(email string) (*User, error)
	FindByUsername(username string) (*User, error)
	Delete(ID ID) error
	Register(newUser CreateUser) (*User, error)
	ValidateUser(user User) error
	ValidateNewUser(newUser CreateUser) error
}

type service struct {
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
) Service {
	return &service{
		Authenticator: a,
		Finder:        f,
		Deleter:       d,
		Registrator:   r,
		Validator:     v,
	}
}
