package message

type Service struct {
	Finder
	Deleter
	Creator
}

// NewService creates a new user service
func NewService(
	r Finder,
	d Deleter,
	c Creator,
) *Service {
	return &Service{
		Finder:  r,
		Deleter: d,
		Creator: c,
	}
}
