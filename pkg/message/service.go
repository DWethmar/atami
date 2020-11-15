package message

// Service defines interations with users
type Service interface {
	FindByID(ID int) (*Message, error)
	Find() ([]*Message, error)
	Delete(ID int) error
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
