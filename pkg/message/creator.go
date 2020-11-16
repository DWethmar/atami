package message

// CreatorRepository defines a messsage listing repository
type CreatorRepository interface {
	Create(newMessage CreateMessage) (*Message, error) // return int
}

// Creator creates messages.
type Creator struct {
	validator  *Validator
	createRepo CreatorRepository
}

// Create a new message
func (m *Creator) Create(newMessage CreateMessage) (*Message, error) {
	return m.createRepo.Create(newMessage)
}

// NewCreator returns a new Listing
func NewCreator(r CreatorRepository) *Creator {
	return &Creator{
		NewDefaultValidator(),
		r,
	}
}
