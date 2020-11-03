package user

// Service defines interations with users
type Service interface {
	SearchByEmail(email string) ([]*User, error)
	ReadOne(ID ID) (*User, error)
	ReadAll() ([]*User, error)
	Delete(ID ID) error
	Create(newUser NewUser) (*User, error)
}

type service struct {
	Searcher
	Reader
	Deleter
	Creator
}

// NewService creates a new user service
func NewService(
	s Searcher,
	r Reader,
	d Deleter,
	c Creator,
) Service {
	return &service{
		Searcher: s,
		Reader:   r,
		Deleter:  d,
		Creator:  c,
	}
}
