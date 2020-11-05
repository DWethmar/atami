package user

// Service defines interations with users
type Service interface {
	FindAll() ([]*User, error)
	FindByID(ID ID) (*User, error)
	FindByEmail(email string) (*User, error)
	FindByUsername(username string) (*User, error)
	Delete(ID ID) error
	Register(newUser NewUser) (*User, error)
}

type service struct {
	Finder
	Deleter
	Registrator
}

// NewService creates a new user service
func NewService(
	f Finder,
	d Deleter,
	r Registrator,
) Service {
	return &service{
		Finder:      f,
		Deleter:     d,
		Registrator: r,
	}
}
