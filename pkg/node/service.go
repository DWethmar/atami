package node

// Service organizes interactions with nodes
type Service struct {
	Finder
	Deleter
	Creator
	Validator
}

// NewService creates a new user service
func NewService(
	r Finder,
	d Deleter,
	c Creator,
) *Service {
	return &Service{
		Finder:    r,
		Deleter:   d,
		Creator:   c,
		Validator: *NewDefaultValidator(),
	}
}
