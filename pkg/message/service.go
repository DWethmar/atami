package message

// Service defines interations with users
type Service interface {
	ReadOne(ID ID) (*Message, error)
	ReadAll() ([]*Message, error)
	Delete(ID ID) error
	Create(newMessage NewMessage) (*Message, error)
}

type service struct {
	Reader
	Deleter
	Creator
}

// NewService creates a new user service
func NewService(
	r Reader,
	d Deleter,
	c Creator,
) Service {
	return &service{
		Reader:  r,
		Deleter: d,
		Creator: c,
	}
}
