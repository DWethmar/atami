package message

// NewMessage model
type NewMessage struct {
	Text            string
	CreatedByUserID int
}

// CreatorRepository defines a messsage listing repository
type CreatorRepository interface {
	Create(newMessage NewMessage) (*Message, error) // return int
	// add find repo and retrieve it.
}

// Creator creates messages.
type Creator struct {
	createRepo CreatorRepository
}

// Create a new message
func (m *Creator) Create(newMessage NewMessage) (*Message, error) {
	return m.createRepo.Create(newMessage)
}

// NewCreator returns a new Listing
func NewCreator(r CreatorRepository) *Creator {
	return &Creator{r}
}
