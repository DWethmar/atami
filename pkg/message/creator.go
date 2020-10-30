package message

// New model
type NewMessage struct {
	Content string
}

// CreatorRepository defines a messsage listing repository
type CreatorRepository interface {
	Create(newMessage New) (Model, error)
}

// Creator creates messages.
type Creator struct {
	createRepo CreatorRepository
}

// Create a new message
func (m *Creator) Create(newMessage New) (Model, error) {
	return m.createRepo.Create(newMessage)
}

// NewCreator returns a new Listing
func NewCreator(r CreatorRepository) *Creator {
	return &Creator{r}
}
