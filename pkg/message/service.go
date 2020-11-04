package message

// Service defines interations with users
type Service interface {
	FindByID(ID ID) (*Message, error)
	FindAll() ([]*Message, error)
	Delete(ID ID) error
	Create(newMessage NewMessage) (*Message, error)
}

type service struct {
	Finder
	Deleter
	Creator
}

// NewService creates a new user service
func NewService(
	r Finder,
	d Deleter,
	c Creator,
) Service {
	return &service{
		Finder:  r,
		Deleter: d,
		Creator: c,
	}
}
