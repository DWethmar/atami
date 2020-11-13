package message

import "github.com/dwethmar/atami/pkg/model"

// NewMessage model
type NewMessage struct {
	Text      string
	CreatedBy model.UserID
}

// CreatorRepository defines a messsage listing repository
type CreatorRepository interface {
	Create(newMessage NewMessage) (*Message, error)
}

// Creator creates messages.
type Creator struct {
	createRepo CreatorRepository
}

// Create a new message
func (m *Creator) Create(newMessage NewMessage) (*model.Message, error) {
	message, err := m.createRepo.Create(newMessage)
	if err != nil {
		return nil, err
	}
	return toMessage(message), nil
}

// NewCreator returns a new Listing
func NewCreator(r CreatorRepository) *Creator {
	return &Creator{r}
}
