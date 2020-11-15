package user

// Service defines interations with users
type Service struct {
	Finder
	Deleter
	Creator
}

// NewService creates a new user service
func NewService(
	f Finder,
	d Deleter,
	r Creator,
) *Service {
	return &Service{
		Finder:  f,
		Deleter: d,
		Creator: r,
	}
}
