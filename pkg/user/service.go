package user

// Service defines interations with users
type Service interface {
	FindAll() ([]*User, error)
	FindByID(ID ID) (*User, error)
	FindByEmail(email string) (*User, error)
	Delete(ID ID) error
	Create(newUser NewUser) (*User, error)
}

type service struct {
	Finder
	Deleter
	Creator
}

// NewService creates a new user service
func NewService(
	f Finder,
	d Deleter,
	c Creator,
) Service {
	return &service{
		Finder:  f,
		Deleter: d,
		Creator: c,
	}
}
